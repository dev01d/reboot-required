# Reboot required

This program checks for the need to reboot on Ubuntu and Red Hat-based systems. If there are any updates applied to the kernel, it lets you know if a reboot is required.

By default it only tells you if a reboot is required or not. The verbose mode displays all packages causing the required reboot.

## Usage

```bash
Usage: rr [OPTIONS]

-h, --help                Show this help
-v, --verbose             Packages causing the need to reboot
    --version             Print version
```

## Install

- Homebrew:

```shell
brew install dev01d/tap/reboot-required
```

- Apt
<!-- /* spellchecker: disable */ -->

```shell
curl -fsSL https://dev01d.fury.site/apt/gpg.key | sudo gpg --dearmor -o /usr/share/keyrings/dev01d.gpg
```

```bash
echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/dev01d.gpg] \
 https://dev01d.fury.site/apt/ * *" \
| sudo tee -a /etc/apt/sources.list.d/dev01d.list > /dev/null
```

```bash
sudo apt-get update; sudo apt install reboot-required
```

- Yum

```shell
sudo echo """\
[fury]
name=dev01d repo
baseurl=https://dev01d.fury.site/yum/
enabled=1
gpgcheck=0
""" > /etc/yum.repos.d/dev01d.repo

yum install reboot-required
```

<!-- /* spellchecker: enable */ -->
