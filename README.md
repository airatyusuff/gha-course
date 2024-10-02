# GHA-course

- OReilly course on GHA: https://learning.oreilly.com/course/learning-github-actions/9781837639137/


Command to create service principal, for role-based access, with a contributor role (access to manage full resources but not admin-level), and scoped to a resource group (to limit to those resources), and with JSON output to grab details for Azure SDK authentication

```
az ad sp create-for-rbac --name <<name>> --role <<role_type>> --scopes <<resource id(s)>> --sdk-auth
```

Note: Didn't work for my bjss account (need root level permissions)