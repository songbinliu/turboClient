# turboClient
A rest client for Turbonomic API server. (Currently it only provides an way to add license file to Turbonomic server.)

# Run it

### Build it
```bash
go build
```

### Set the host, login credentials of the turbonomic.server
```terminal
Usage of ./turboClient:
  -fname string
    	the xml license file. (default "./data/license.xml")
  -host string
    	the address of turbo.server (default "https://localhost:9400")
  -logtostderr
    	log to standard error instead of files
  -pass string
    	password to login to turbo.server (default "a")
  -user string
    	username to login to turbo.server (default "administrator")
  -v value
    	log level for V logs

```

### Run it
```bash
./turboClient -v=3 --logtostderr --host=https://10.10.174.134 --user=administrator --pass=a --fname=./data/trial.license.xml 
```

If the license file is validate, it will print some log like this:
```terminal
I1107 09:57:09.112003    1107 main.go:27] result={"licenseOwner":"songbinliu","email":"songbin.liu@turbonomic.com","expirationDate":"Dec 10 2017","features":["action_script","active_directory","aggregation","applications","automated_actions","cloud_targets","custom_reports","customized_views","deploy","fabric","full_policy","group_editor","historical_data","loadbalancer","multiple_vc","optimizer","scoped_user_view","storage","trial","vmturbo_api"],"numSocketsLicensed":100,"numSocketsInUse":12,"isValid":true}
````

However, if the license file is not valide, then it will be like:
```terminal
I1107 09:59:39.743273    1113 main.go:27] result={"isValid":false}
```

