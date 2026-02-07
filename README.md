# Sential-host
Sential-host is an EDR style, local, open-source security agent for linux systems. we are only at our first commit. its pretty solid.

# WARNING only run in safe mode. this is an early version.
this will only log the malicious traffic but not stop it. we need to test unsafe mode further and develop it further as not safe mode can destroy networking. I do not take accountability for your incompetence with this tool.

it stays between the user networking and the kernel and reviews if your traffic makes sense. if it doesnt. it shuts it down. prevents cyberattacks as they happen.

# I focused on.
- early detection of commodity malware and abuse
- behavioural detection

# Threat model.
This logs malicious traffic from skids and hackers. its not going to stop full government survailence but it works for logging some black hats.

# LICENSE

This project is licensed under the MIT License.  
See the [LICENSE](LICENSE) file for details.

# Credits
wet-cat(me)
