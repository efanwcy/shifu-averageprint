



# shifu-averageprint
基于shifu项目下，定期获取酶标仪的get_measurement接口，每五分钟打印一次平均值。
main.go为求接口值的代码，文件夹中包含.yaml配置文件，，docker文件与dockerfile


步骤1.跟随shifu官方文档部署Docker Desktop，这里我使用macos系统部署。
<img width="491" alt="image" src="https://github.com/efanwcy/shifu-averageprint/assets/101723328/444ed130-914e-43b1-86cb-73f6ccd96d61">



步骤2：配置shifu所需依赖项
golang 官方文档部署 https://go.dev/dl/
Docker，按照Docker Desktop部署。
kubectl：按照kubernetes官网部署，这里我选择的macos版本https://kubernetes.io/docs/tasks/tools/
kind  按照官网部署https://kind.sigs.k8s.io/docs/user/quick-start/
kubebuilder的部署：https://github.com/kubernetes-sigs/kubebuilder
步骤2 部署shifu
打开Docker Desktop的kubernetes。 ![dc9f1db3d750d87895ecf806e58d504](https://github.com/efanwcy/shifu-averageprint/assets/101723328/1f3f20ca-9f85-4aee-bf96-029245427da4)

输入sudo docker ps ![38eb5b6328c7ae6b3f7abc36b18ebf3](https://github.com/efanwcy/shifu-averageprint/assets/101723328/9b0e757e-feb6-4ba7-bbb8-35f1512171e6)

使用curl -sfL https://raw.githubusercontent.com/Edgenesis/shifu/main/test/scripts/shifu-demo-install.sh | sudo sh -命令部署shifu ![3c94553bf0655f9204b1ea1d273160f](https://github.com/efanwcy/shifu-averageprint/assets/101723328/d3b67456-7d4e-4395-8f7b-eca79467cae5)

cd命令进入shifudemos，使用sudo kubectl run --image=nginx:1.21 nginx命令启动nginx,使用sudo kubectl get pods -A | grep nginx命令去查看nginx ![d41ed4aede70cede8f3e78aae374107](https://github.com/efanwcy/shifu-averageprint/assets/101723328/6dd46bd9-d1c4-4fb9-9a07-69c9dd0bdfff)


步骤3 与酶标仪交互，使用sudo kubectl apply -f run_dir/shifu/demo_device/edgedevice-plate-reader命令启动酶标仪，并通过sudo kubectl get pods -A|grep plate去查看 ![82fd9024d67f6463c08569c196c27e2](https://github.com/efanwcy/shifu-averageprint/assets/101723328/1413d71a-4c4b-4054-b888-864205d986bd)
 
sudo kubectl exec -it nginx -- bash进入nginx,输入curl "deviceshifu-plate-reader.deviceshifu.svc.cluster.local/get_measurement"命令可以得到 ![8d61049da9d46502b02b20bbe455165](https://github.com/efanwcy/shifu-averageprint/assets/101723328/e2cddbb3-1a54-4f3d-92f3-9ceb4424ed0f)



步骤4 将main.go文件配置docker file，并使用docker build -t averageprint .构建Docker镜像。
docker login在本地登录Docker Hub，docker tag averageprint golangtrainee/average标记镜像，并使用docker push golangtrainee/averageprint推送到Dockerhub <img width="688" alt="1053d5f4be444b42f47040a3c9614b3" src="https://github.com/efanwcy/shifu-averageprint/assets/101723328/a31a3102-5228-4c6e-9b22-d1b508bb1474">

编写averageprint-deployment.yaml并运行。  ![cfd0b0e6cd28cc4f69649402593d780](https://github.com/efanwcy/shifu-averageprint/assets/101723328/0c284eea-fb86-402f-9069-22986d0c92af)

执行kubectl logs -f deployment/averageprint-deployment命令查看日志结果   ![c871da94f0f730dd145f83bd77c5b83](https://github.com/efanwcy/shifu-averageprint/assets/101723328/6f089f25-4944-4f33-bf58-9ffa7e0044c1)
