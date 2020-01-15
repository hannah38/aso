/*
MIT License

Copyright (c) Microsoft Corporation. All rights reserved.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE
*/

package storages

import (
	"context"
	"errors"
	"log"

	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2019-04-01/storage"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
)

type azureStorageManager struct{}

func getStoragesClient() storage.AccountsClient {
	storagesClient := storage.NewAccountsClient(config.SubscriptionID())
	a, err := iam.GetResourceManagementAuthorizer()
	if err != nil {
		log.Printf("failed to initialize authorizer: %v\n", err)
	}
	storagesClient.Authorizer = a
	storagesClient.AddToUserAgent(config.UserAgent())
	return storagesClient
}

// CreateStorage creates a new storage account
func (_ *azureStorageManager) CreateStorage(ctx context.Context, groupName string,
	storageAccountName string,
	location string,
	sku azurev1alpha1.StorageSku,
	kind azurev1alpha1.StorageKind,
	tags map[string]*string,
	accessTier azurev1alpha1.StorageAccessTier,
	enableHTTPsTrafficOnly *bool, dataLakeEnabled *bool) (result storage.Account, err error) {

	storagesClient := getStoragesClient()

	//Check if name is available
	storageType := "Microsoft.Storage/storageAccounts"
	checkAccountParams := storage.AccountCheckNameAvailabilityParameters{Name: &storageAccountName, Type: &storageType}
	checkNameResult, err := storagesClient.CheckNameAvailability(ctx, checkAccountParams)
	if err != nil {
		return result, err
	}
	if dataLakeEnabled == to.BoolPtr(true) && kind != "StorageV2" {
		log.Printf("Cannot create storage account. Datalake enabled storage account must be of kind: StorageV2")
		return result, errors.New("unable to create datalake enabled storage account")
	}
	if *checkNameResult.NameAvailable == false {
		log.Println("storage account not available: " + checkNameResult.Reason)
		if checkNameResult.Reason == storage.AccountNameInvalid {
			return result, errors.New("AccountNameInvalid")
		} else if checkNameResult.Reason == storage.AlreadyExists {
			return result, errors.New("AlreadyExists")
		}
	}

	sSku := storage.Sku{Name: storage.SkuName(sku.Name)}
	sKind := storage.Kind(kind)
	sAccessTier := storage.AccessTier(accessTier)

	params := storage.AccountCreateParameters{
		Location: to.StringPtr(location),
		Sku:      &sSku,
		Kind:     sKind,
		Tags:     tags,
		Identity: nil,
		AccountPropertiesCreateParameters: &storage.AccountPropertiesCreateParameters{
			AccessTier:             sAccessTier,
			EnableHTTPSTrafficOnly: enableHTTPsTrafficOnly,
			IsHnsEnabled:           dataLakeEnabled,
		},
	}

	//log.Println(fmt.Sprintf("creating storage '%s' in resource group '%s' and location: %v", storageAccountName, groupName, location))
	future, err := storagesClient.Create(ctx, groupName, storageAccountName, params)
	if err != nil {
		return result, err
	}

	return future.Result(storagesClient)

}

// Get gets the description of the specified storage account.
// Parameters:
// resourceGroupName - name of the resource group within the azure subscription.
// accountName - the name of the storage account
func (_ *azureStorageManager) GetStorage(ctx context.Context, resourceGroupName string, accountName string) (result storage.Account, err error) {
	storagesClient := getStoragesClient()
	return storagesClient.GetProperties(ctx, resourceGroupName, accountName, "")
}

// DeleteStorage removes the resource group named by env var
func (_ *azureStorageManager) DeleteStorage(ctx context.Context, groupName string, storageAccountName string) (result autorest.Response, err error) {
	storagesClient := getStoragesClient()
	return storagesClient.Delete(ctx, groupName, storageAccountName)
}