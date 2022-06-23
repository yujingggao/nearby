# nearby

A social network web app that allows users to create and list personal posts, browse and search posts within a distance

## Tech Stack: 
-   Frontend: React JS, JWT
-   Backend: Go, Elasticsearch
-   REST API: 
-   Cloud: Google Cloud(Google App Engine)

## Feature

## Requirements
### Google Compute Engine (GCE)
We will be using `GCE VM` to develop this `Go` program.  Create a `Firewall rule` in Google Cloud, and see the [GCE documentation](https://cloud.google.com/compute/docs)here to setup a GCE instance

### VSCode Config

#### Create SSH Key
We will use `VSCode` for development.  Run the following command in your terminal/git bash to create `SSH Key` for your `GCE instance`
```shell
>$ ssh-keygen -t rsa -f ~/.ssh/gcekey -C <YOUR_USERNAME>
```

Verify that both public and private keys are created
```shell
>$ ls ~/.ssh/
```

Use the following command to print the public key content and verify that the content is in the format of “ssh-rsa KEY_VALUE YOUR_USERNAME”
```shell
>$ cat ~/.ssh/gcekey.pub
```

Copy the public key and paste it to your `VM instance`

#### Set up Remote SSH in IDE
In `VSCode`, install the `Remote SSH` extension, navigate to the `Remote Explorer` window, click the plus button to create a new target.

In the popup window, enter the following command
```shell
>$ ssh -i ~/.ssh/gcekey <YOUR_USRENAME>@<YOUR_GCE_EXTERNAL_IP_ADDRESS>
```

In the next window to update SSH configuration file, select the one under `~/.ssh/config`, and verify the new target is created in the window 

#### SSH Key and Config Permissions
Open your terminal and navigate to the `~/.ssh` directory and update the access permissions for `SSH Key` and `Configuration`
```shell
>$ chmod 600 config
>$ chmod 600 gcekey
```

#### Connect to GCE Instance in VSCode
Under the `SSH Targets` window, right-click your target and select `Connect to Host in Current Window`, and verify if the SSH status at the bottom-left of your VSCode window is green.

Open the `Explorer` window, and open a directory.  This will be your working directory in your `VM`.  

## Install Go
Install `Go` language support plugin from the VSCode Extension.

Install `Go` on your GCE instance by running the following command in your terminal window
```shell
>$ sudo add-apt-repository ppa:longsleep/golang-backports  
>$ sudo apt-get update  
>$ sudo apt-get install golang-go
```

Verify if the installation is completed
```shell
>$ go version
```

# Run server
```shell
>$ go run main.go
```

# Elasticsearch
`Elasticsearch` is an open-source, distributed, RESTful search engine. As the heart of the Elastic Stack, it centrally stores your data so you can query the data quickly.  This project uses Elasticsearch as its `Database`

Some basic concepts in `Elasticsearch` includes:
-   Index: like a database in a relational database management system.
-   Type: like a table of a database, but deprecated after Elasticsearch 6.0.
-   Document: like a row of a table.
-   Property: like a column of a table.
-   Mapping: like a schema of a database table.

## Set up Elasticsearch on GCE instance
`SSH` to your `GCE instance` and enter the following commands to install `Java`
```shell
>$ sudo apt install default-jre
```

Verify `Java` is intalled correctly
```shell
>$ sudo java -version
```

Install `Elasticsearch` on your GCE instance
```shell
>$ sudo apt install apt-transport-https

>$ wget -qO - https://artifacts.elastic.co/GPG-KEY-elasticsearch | sudo apt-key add -

>$ sudo sh -c 'echo "deb https://artifacts.elastic.co/packages/7.x/apt stable main" > /etc/apt/sources.list.d/elastic-7.x.list'

>$ sudo apt update

>$ sudo apt install elasticsearch

```

Open the `configuration` file
```shell
>$ sudo vim /etc/elasticsearch/elasticsearch.yml
```

Make the following changes:
1. Under the `Network` section, Uncomment `network.host`(line 56) and `http.port`(line 60) by deleting the leading “#”. Set network.host to 0.0.0.0, http.port to 9200.
```shell
network.host: 0.0.0.0
http.port: 9200
```

2. In the `Discovery` section, add the following line above `discovery.seed_hosts` line since we are running `Elasticsearch` in a `single node` mode
```shell
discovery.type: single-node
```

3. In the `Various` section, below `action.destructive` line, add the following line to ennable the security feature of `Elasticsearch`
```shell
discovery.type: single-node
```

Save the changes and verify
```shell
>$ sudo cat /etc/elasticsearch/elasticsearch.yml|grep "^[^#;]"
```

Make `Elasticsearch` auto-start every time we start our `GCE instance`
```shell
>$ sudo systemctl enable elasticsearch
```

## Start Elasticsearch
```shell
>$ sudo systemctl start elasticsearch
```

Check the status of Elasticsearch.  Type `q` to exit the status page
```shell
>$ sudo systemctl status elasticsearch
```

## Update Elasticsearch password
Add a new user
```shell
>$ sudo /usr/share/elasticsearch/bin/elasticsearch-users useradd <YOUR_NEW_USER_NAME> -p <YOUR_NEW_USER_PASSWORD> -r superuser
```

Open a browswer, enter `http://YOUR_GCP_INSTANCE_EXTERNAL_IP:9200`, and login with the username and password you created

## Connect Go project to Elasticsearch
### Create Elasticsearch Indexes
Install the Elasticsearch library in the same directory of your `Go` Project
```shell
>$ cd go/src/nearby
>$ go get github.com/olivere/elastic/v7
```

### Test Elasticsearch index creation
Start the program with `go run main.go`.  If it returns with `Indexes index is created` it's successful. Use `Ctrl` + `C` to stop.

Verify Post index:  `http://YOUR_GCP_INSTANCE_EXTERNAL_IP:9200/post?pretty` 
Verify User index:  `http://YOUR_GCP_INSTANCE_EXTERNAL_IP:9200/user?pretty`

# Google Cloud Storage (GCS)
`GCS` is a Powerful, Simple and Cost-effective Object Storage Service. Same as `S3` service on AWS.  It behaves like a `file system`, has good availability and durability, is affordable and provides CDN service to serve files in edge servers to reduce loading latency.

Compared to storing media files in a database, storing them in a file system directly is much faster.  Media files increase the size of the database drastically, making it hard for maintenance.  There is also no way to do datase related optimization (like `indexing`) based on a binary file.  

## GCS Buckets
-   Buckets are the basic containers that hold your data. You can image buckets like the root directory in your file system. You can upload any file as an object into your bucket.
-   The bucket name must be globally unique, there’s only one single namespace in GCS.
-   You can apply labels to buckets to group buckets together.
-   Objects are the individual pieces of data that you store in GCS. You can image an object like a file in your file system. You can put multiple objects in a single bucket.
-   Objects are immutable, so you can’t make incremental changes, but only overwrite the whole object.

Refer to [this doc](https://cloud.google.com/storage/docs/creating-buckets) to create your GCS bucket. 

test
