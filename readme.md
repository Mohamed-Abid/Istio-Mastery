# Deploy an E-commerce Application using Istio


## Prerequisites

- [Minikube](https://minikube.sigs.k8s.io/docs/start/) installed and running
- [kubectl](https://kubernetes.io/docs/tasks/tools/) installed and configured
- [Istio](https://istio.io/latest/docs/setup/getting-started/) installed in your Minikube cluster

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/Mohamed-Abid/Istio-Mastery.git
   cd Istio-Mastery
    ```

2. Start Minikube:
   ```bash
    minikube start
    ```

3. Enable Istio injection for default namespace:
    ```bash
    minikube addons enable istio
    ```

4. Add any necessary Istio addons:
   ```bash
    istioctl install --set profile=default
    ```
    
5. Enable Istio injection for default namespace:
    ```bash
    kubectl label namespace default istio-injection=enabled
    ```

6. Deploy your deployments:
    ```bash
    kubectl apply -f Manifests/fulldeployment.yaml
    ```

7. Deploy necessary Addons:
    ```bash
    kubectl apply -f Addons
    ```

8. Deploy Istio files:
    ```bash
    kubectl apply -f Manifests/Security.yaml
    kubectl apply -f Manifests/traffic.yaml
    ```

9.  Open the application in a browser by running:
    ```bash
    kubectl port-forward -n istio-system svc/istio-ingressgateway 8080:80
    ```
    
10. Visualize Traffic with Kiali:
    ```bash
    istioctl dashboard kiali
    ```

11. Visualize Metrics with Grafana:
    ```bash
    istioctl dashboard grafana
    ```

