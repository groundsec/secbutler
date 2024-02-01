# Tests

This folder contains files fore **general purpose** testing (not code testing).

## Install tools

The `tools` command dynamically generates a `install_tools.sh` script that installs a series of cybersec tools. The script checks if the requirements are met and then, if every requirement is met, it proceeds on installing the tools. Since the installation of the requirements is out of scope a Docker image with the requirements pre-installed is provided for quick testing.

Build the Docker image:

```bash
docker build . -t groundsec/secbutler-tools
```

Then run `secbutler tools` to generate the `install_tools.sh` script.
Go to `$HOME/.secbutler` and run:

```bash
docker run -it --rm -v $PWD:/secbutler groundsec/secbutler-tools
```

Then, inside the container, run:

```bash
/secbutler/install_tools.sh
```
