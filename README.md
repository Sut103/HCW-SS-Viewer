# HCW-SS-Viewer
## 開発コンテナ上での実行
* 以下、VSCodeのdevcontainerを想定

### DynamoDBテーブルを作成
* `http://localhost:8001`にてGUIで設定可能
* テーブル名は任意
* 現状、PKとして`URL`があれば良い

### エミュレート時の環境変数を指定
```plaintext:./local_invoke/local_invoke_env.json
{
    "Parameters": {
        "DYNAMO_ENDPOINT":"http://dynamodb-local:8000",
        "DYNAMO_TABLE_NAME": YOUR_TABLE_NAME
    }
}
```

### sam build
* ルートディレクトリにて

```shell
sam build
```

### sam local
* Lambdaの実行
```
sam local invoke --container-host host.docker.internal --container-host-interface host.docker.internal --docker-network hcw-ss-viewer_devcontainer_default -n local_invoke/local_invoke_env.json -e local_invoke/get_event.json
```
* API Gatewayの実行
```
sam local start-api --warm-containers LAZY --container-host host.docker.internal --container-host-interface host.docker.internal --docker-network hcw-ss-viewer_devcontainer_default -n local_invoke/local_invoke_env.json
```