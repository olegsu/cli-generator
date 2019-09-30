# CLI-GENERATOR
Generate CLI entrypoints from spec file
Any feedback on the spec is welcome

## Run Example
* `git clone git@github.com:olegsu/cli-generator.git`
* `cd cli-generator`
* `make example`


The main idea is to define a spec that can describe a CLI in YAML or JSON format.
Spec must include:
1. Name of the application(project)
    * Example: `docker`, `kubectl`
2. Global flags the CLI would support
    * Example: `kubectl --context`
3. Commands, including nasted sub-commands and aliases
    * Example: `kubectl get`, `docker system prune`, `kubectl get po`, `kubectl get pods`
4. Flags per command, including types [array of strings] [enum] definition, default values and environment variables are eqvivalent 
    * Example: `kubectl logs -f [NAME]`
5. Positional argumens, including multiple arguments
    * Example: `kubectl delete po [PO_1] [PO_2]`
6. Auto generate help or provide a way to override help per command
