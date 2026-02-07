# Sential-host
Sential-host is an EDR style, local, open-source security agent for linux systems. we are only at our first commit. its pretty solid.

# WARNING only run in safe mode. this is an early version.
this will only log the malicious traffic but not stop it. we need to test unsafe mode further and develop it further as not safe mode can destroy networking. I do not take accountability for your incompetence with this tool. You are responsable for any damages done. USE SAFE MODE.

# installation
1 
sudo apt update
sudo apt install -y golang iptables libnetfilter-queue-dev

2

''git clone https://github.com/wet-cat/Sential-host.git

cd Sential-host

go mod tidy''

3

go build -o sentinel ./cmd/sentinel

4

// set rules

sudo ./scripts/setup_iptables.sh

// if you want to remove these rules. do this

sudo iptables -D INPUT -j NFQUEUE --queue-num 0

sudo iptables -D OUTPUT -j NFQUEUE --queue-num 0

5

sudo ./sentinel -safe=true


to close it do this:

press ctrl + c then run:

sudo iptables -D INPUT -j NFQUEUE --queue-num 0

sudo iptables -D OUTPUT -j NFQUEUE --queue-num 0


# I focused on.
- early detection of commodity malware and abuse
- behavioural detection

# Threat model.
This logs malicious traffic from skids and hackers. its not going to stop full government survailence but it works for logging some black hats.

# LICENSE

This project is licensed under the MIT License.  
See the [LICENSE](LICENSE) file for details.

# Contributors
wet-cat(me)
Open for more!


# TODO.
- make unsafe mode work without potential damages.
- block malicious ip's successfully without networking damages.
- 
