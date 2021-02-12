# To Start App

```bash
docker-compose up
```



# Playground

Found at `http://localhost:8080/`



# Queries, Mutations

### Create User

```
mutation {
  createUser(input: {username: "cooki", password:"insabgho123"}){
    username
    password
  }
}
```



### Create Link

```
mutation {
  createLink(input: {title: "new link", address:"http://address.org"}){
    title,
    user{
      username
    }
    address
  }
}
```



### Get Links

```
query {
	links{
    title
    address,
    user{
      username
      password
    }
  }
}
```



### Get Users

```
query {
	users{
		username,
    password
  }
}
```

