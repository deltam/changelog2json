# changelog2json

Convert [changelog memo](http://0xcc.net/unimag/1/) to JSON.

Suggest to use with [jq](https://stedolan.github.io/jq/).

# Install

```sh
$ go get -u github.com/deltam/changelog2json
```

# Usage

sample changelog

```sh
$ cat changelog.txt
2018-06-29  Nanashi Gombeh  <deltam@gmail.com>

        * life: shopping list
        - milk
        - eggs

        * todo: dev
        - write test code
        - refactoring

2018-06-28  Nanashi Gombeh  <deltam@gmail.com>

        * lunch: Ramen Shop Hogehoge
        Very good!

```

convert to json(and pipe jq).

```sh
$ changelog2json changelog.txt | jq
[
	{
		"date": "2018-06-29",
		"name": "Nanashi Gombeh",
		"mail": "deltam@gmail.com",
		"tags": "life",
		"title": "shopping list",
		"body": "        - milk\n        - eggs\n\n",
		"line": 3
	},
	{
		"date": "2018-06-29",
		"name": "Nanashi Gombeh",
		"mail": "deltam@gmail.com",
		"tags": "todo",
		"title": "dev",
		"body": "        - write test code\n        - refactoring\n\n",
		"line": 7
	},
	{
		"date": "2018-06-28",
		"name": "Nanashi Gombeh",
		"mail": "deltam@gmail.com",
		"tags": "jj",
		"title": "Ramen Shop Hogehoge",
		"body": "        Very good!\n",
		"line": 13
	}
]
```

Filtering by tag

```sh
$ changelog2json changelog.txt | jq '.[] | select(.tags == "todo")'
[
	{
		"date": "2018-06-29",
		"name": "Nanashi Gombeh",
		"mail": "deltam@gmail.com",
		"tags": "todo",
		"title": "dev",
		"body": "        - write test code\n        - refactoring\n\n",
		"line": 7
	}
]
```

more filtering tips, to see [jq manual](https://stedolan.github.io/jq/manual/).

# License

MIT
