
## Aliases

```BASH
alias k=kubectl
alias kdy="kubectl --dry-run=client -o yaml"
```

context
```BASH
k config get-contexts
k config use-context <context-name>
```

Get pods from all namespaces:
`k get po -A`

```BASH
kdy run po pod-name --image=some-image > somefile.yaml
k apply -f somefile.yaml

sudo vim /etc/hosts
```

rollback:
`k rollout undo deployment/<deployment-name>`

cpu usage (pods):<br>
`k top po`

TODO: logs command<br>
`k logs <podname>`

Abbreviations:

| Abbreviation | full name       |
| ------------ | --------------- |
| svc          | service         |
| sa           | service-account |
| deploy       | deployment      |
| po           | pod             |
| cm           | config-map      |
| ns           | namespace       |
| cj           | cronjob         |
