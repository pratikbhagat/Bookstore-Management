# Bookstore-Management
Bookstore-Management API is a RESTful API that allows users to manage books in a bookstore. It provides endpoints for creating, reading, updating, and deleting books from a database. Users can also retrieve a list of all books, search for books by title, author, or publication, and sort the list of books by various criteria. The API is built using the Go programming language and the MUX framework, and it uses the GORM library to interact with a MySQL database. The API is designed to be easy to use and extend, with clear and concise documentation and a simple, intuitive API design.

![Project Structure](https://github.com/pratikbhagat/Bookstore-Management/blob/main/project-structure-book-management.png)
![Routes](https://github.com/pratikbhagat/Bookstore-Management/blob/main/routes-book-management.png)

## Languages and Tools

<a href="https://code.visualstudio.com/" target="_blank"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/vscode/vscode-original.svg" alt="Visual Studio Code" width="40" height="40"/> </a>
<a href="https://golang.org/" target="_blank"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/go/go-original-wordmark.svg" alt="Go" width="40" height="40"/> </a>
<a href="https://www.mysql.com/" target="_blank"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/mysql/mysql-original-wordmark.svg" alt="MySQL" width="40" height="40"/> </a>
<a href="https://www.docker.com/" target="_blank"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/docker/docker-original-wordmark.svg" alt="Docker" width="40" height="40"/> </a> 
<a href="https://www.postman.com/" target="_blank"> <img src="https://www.vectorlogo.zone/logos/getpostman/getpostman-icon.svg" alt="postman" width="40" height="40"/> </a>
<a href="https://www.linux.org/" target="_blank"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/linux/linux-original.svg" alt="linux" width="40" height="40"/> </a>
<a href="https://github.com/" target="_blank"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/github/github-original-wordmark.svg" alt="GitHub" width="40" height="40"/> </a>

## Getting Started
<ul>
  <li><b>Prerequisites</b>
    <ul>
      <li>Golang</li>
      <li>MySQL</li>
    </ul>
  </li>
  <li><b>Installation</b>
    <ul>
      <li>Change Directory to <b><i>$GOROOT/src/</i></b></li>
      <li>Clone repository into the directory.
        <ul>
          <li><code>git clone https://github.com/xx/Bookstore-Management.git</code></li>
        </ul>
      </li>
    </ul>
  </li>
  <li><b>Running service</b>
    <ul>
      <li><b>Running service using service and MakeFile.</b>
        <ul>
          <li>Edit bookstore-management.service file.</li>
          <li><i>You need to update “ExecStart” in service file to project-directory-relative-path/main</i></li>
          <li><code>sudo --preserve-env=PATH env make all</code></li>
        </ul>
      </li>
      <li><b>Stopping Using service file</b>
        <ul>
          <li><code>sudo systemctl stop bookstore-management.service</code></li>
        </ul>
      </li>
      <li><b>Running using Docker</b>
        <ul>
          <li><code>sudo docker build -t schema-service:latest -f deployer/docker_deployer/Dockerfile .</code></li>
          <li><code>sudo docker run -p 9080:9080 -v /var/log/apps/schema-service/apiserver.log:/var/log/apps/schema-service/apiserver.log -v /etc/schema-service.yaml:/etc/schema-service.yaml schema-service:latest</code></li>
        </ul>
      </li>
      <li><b>Stopping using Docker</b>
        <ul>
          <li><code>sudo docker stop CONTAINER</code></li>
        </ul>
      </li>
    </ul>
  </li>
  <li><b>Test the service</b>
    <ul>
      <li><code>http://localhost:9080/health</code></li>
    </ul>
  </li>
  <li><b>Start Making API calls</b>
  <ul>
      <li>Use the postman collection (Book Management.postman_collection.json)</li>
    </ul>
  </li>
</ul>



