# tatsumaki-go
A Go wrapper for Tatsumaki's API, one of the most popular bots on Discord.

The current version is 1.0.0-release.

## Tutorial

First, create an instance of the Tatsumaki client which will be
your interface to interact with the API.

Then, call your desired endpoint methods on the client.

```go
// Create client.
tatsumakiClient = tatsumakigo.New("YOUR API TOKEN")

// Get a user.
user, err := tatsumakiClient.User("User ID")

if err == nil {
	// Do stuff with the user response here.
	fmt.Println("User's required XP until next level up: ", user.LevelProgress.RequiredXp)
}
```

## License
Copyright &copy;2019 Hassie.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.