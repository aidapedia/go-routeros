
<h1 align="center">
  <br>
  Go-RouterOS
  <br>

</h1>

<h4 align="center">A minimalist wrapper that contain implementation of <a href="https://github.com/go-routeros/routeros">router-os</a>.</h4>

<p align="center">
<a href="http://makeapullrequest.com">
    <img src="https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=shields"
         alt="PR">
  </a>
</p>

<p align="center">
  <a href="#keys-features">Keys Features</a> •
  <a href="#how-to-use">How To Use</a> •
  <a href="#roadmap">Roadmap</a> •
  <a href="#license">License</a> •
  <a href="#support">Support</a>
</p>

## Key Features
* Easy to add more module.
* Know your data structure for each module.

## How to Use

```
    routerBuilder := routeros.NewRouterOS(&routeros.Options{
		Address:  "<your ip>:8728",
		Username: "<your user-name>",
		Password: "<your password>",
	})
	err := routerBuilder.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer routerBuilder.Close()

	active, err := driver.New(routerBuilder, module.HotspotModule)
	if err != nil {
		log.Fatal(err)
	}

	// Example of using Print
	activeRes, err := active.Print(context.Background(), model.PrintRequest{
		Where: []model.Where{
			{
			 	Field:    "address",
			 	Value:    "10.10.10.53",
			 	Operator: model.OperatorEqual,
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	for _, a := range activeRes {
		fmt.Println(a)
	}
```


## Roadmap
We have a lot of things to provide all API on this SDK. Currently, I work by my priority like this.

- [ ] IP
- [ ] System


## License
The MIT License (MIT) 2017 - [Kurniaji Gunawan](https://github.com/kurniajigunawan/). Please have a look at the [LICENSE.md](LICENSE.md) for more details.

## Support
If you love this respository, don't forge to support me.

[<img style="width: 200px;" src="https://i.ibb.co.com/mCsbv8NQ/378162407-9b80104f-ee5b-4a13-8030-ef9a87ed9836.png">](https://buymeacoffee.com/aidapedia)
