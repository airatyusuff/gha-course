
param prefix string = 'ghacourse'
param location string = resourceGroup().location
param app_name string = 'ghawebapp'

var storageAccountName = '${prefix}stgacc'

// create a resource
resource storage_account 'Microsoft.Storage/storageAccounts@2023-05-01' = {
  name: storageAccountName
  location: location
  sku: {
    name: 'Standard_LRS'
  }
  kind: 'StorageV2'
}

resource container_registry 'Microsoft.ContainerRegistry/registries@2023-11-01-preview' = {
  name: '${prefix}acrgha'
  location: location
  sku: {
    name: 'Basic'
  }
  properties: {
    adminUserEnabled: true
  }
}

resource service_bus 'Microsoft.ServiceBus/namespaces@2023-01-01-preview' = {
  name: '${prefix}sbgha'
  location: location
}

resource hostingPlan 'Microsoft.Web/serverfarms@2023-12-01' = {
  name: 'asp-${app_name}'
  location: location
  kind: 'linux'
  sku: {
    tier: 'Basic'
    name: 'B1'
    size: 'B1'
    family: 'B'
    capacity: 1
  }
}

resource name_resource 'Microsoft.Web/sites@2023-12-01' = {
  name: app_name
  location: location
  tags: null
  properties: {
    serverFarmId: hostingPlan.id
    clientAffinityEnabled: false
    httpsOnly: true
  }
  dependsOn: []
}
