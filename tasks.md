# Tasks

## blue-green deployment

1. make new green deployment (e.g. by copying blue deployment manifest)
2. update labels in metadata section
3. verify it is working
4. update the service to point to green deployment (labels)
5. curl worker node, ensure instances of only green deployment is returned

## canary deployment

1. Create a canary Deployment setup for existing deployment, by copy manifest file for existing deployment
2. ensure replica counts are correct in both deployments
3. potentially remove env in svc, only keeping app name
4. curl worker node, ensure instances of both deployments are returned
