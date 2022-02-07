# Echo Api Template

## <font color="cyan">Application Stack</font>
- Golang v1.17.5
- Echo framework v4.6.3
- Struct validator v10.10.0
- Ent orm v0.10.0
- Postgres pg v1.10.4
- Viper environment v1.10.1
- Swag swagger v1.7.9
- Air live reload v1.27.10

### <font color="cyan">Windows users can use git bash to run air. </font>
- curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

## Application Directory Structure
````
├───app  - Application directory
│   └───rest
├───bin  - Air live reload config directory 
│   └───tmp
├───cmd  - Application main directory
├───docs - Swagger documentation directory
├───ent  - Ent orm directory
│   ├───enttest
│   ├───hook
│   ├───migrate
│   ├───predicate
│   ├───runtime
│   ├───schema - Ent schema directory
│   └───user
├───internal - Application core directory
│   ├───controller - Controller directory
│   │   └───user 
│   ├───dto - Data transfer object(dto) directory
│   ├───handle - Application handle directory
│   ├───middleware - Application middleware directory
│   ├───model - Application model(entity) directory
│   ├───response - Application response directory
│   ├───secret - Application secret directory
│   └───validation - Application validation directory
├───pkg - Application package directory 
│   └───config

````

## Swagger generates
![](https://user-images.githubusercontent.com/55887187/152854587-672fd01b-b16f-4705-b2a2-f175b283b5ce.png)
