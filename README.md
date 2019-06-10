# Darklist

## Warning
Darklist is currently in ALPHA stages of development and may not be fully functional. When using this project at this current point in time expect issues.

## Prerequisites
------
Before using Darklist ensure the system running Darklist meets the following requirements:  
1. Minimum Docker version: 18.09.4 or higher  
2. Docker-Compose version 1.18.0 or higher  
3. Docker-Py version 2.6.1 or higher  
4. CPython version: 2.7.13 or higher  
5. OpenSSL version: 1.0.1t or higher  
6. Git version 1.8.3.1 or higher
  
## Install
------
To install Darklist please follow the steps below:  
1. Clone the repository:  
``` cd /tmp && git clone https://github.com/h0lysp4nk/Darklist.git && mv /tmp/Darklist /opt/darklist ```
2. Create a configuration at:  
``` /opt/darklist/volumes/darklist_blacklist/configs/config.json ```  
With the contents of:  
```
{
    "smtp": {
        "host": "192.168.1.87",
        "port": "587",
        "username": "darklist@example.com",
        "password": "mypassword123",
        "ssltls": false,
        "starttls": true
    },
    "runat": "09:00:00"
}
```  
and modify it. Furthermore create another configuration file at:  
``` /opt/darklist/volumes/darklist_blacklist/configs/blacklist.json```  
With the contents of:
```
{
    "blacklists": [
        "bl.spamcop.net",
        "all.spamrats.com",
        "spam.dnsbl.sorbs.net",
        "zen.spamhaus.org",
        "dnsbl.sorbs.net",
        "b.barracudacentral.org"
    ]
}
```  
you can additionally add more or remove blacklists although, this isn't supported!
3. Once you've edited the configuration type the following command to install Darklist:  
``` cd /opt/darklist && bash darklist.sh install ```

## Bug reporting
------
If you're experiencing any issues with Darklist please open an issue under this repository detailing the issue. Please ensure that you provide details about your environment such as operating system, domain, kernal and docker versions etc. If your issue is missing details it will either be put on hold or be closed.

