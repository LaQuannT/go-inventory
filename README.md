# Go Inventory

Go Inventory is a CLI tool that allows users to perform stock keeping.

## Table of Contents

- [Install](#install)
- [Item](#item)
- [Usage](#usage)

<a id="install"></a>
## Install

To use Go-Inventory a few steps must be taken first set up the command for use.

### Steps
- [ ] Fork & Clone Repository
- [ ] Create a .env file with your postgres database variables
- [ ] Build & Install binary

Once repository is forked clone the repository on to your machine:

```
git clone [URI to Repository]
```
Next create a .env file and add the variables below with your postgres database information:

```
touch .env
```
**Variables**
- *PG_USERNAME*
- *PG_PASSWORD*
- *PG_SERVER*
- *PG_PORT*
- *PG_DATABASE*
- *PG_SSLMODE*

Finaly run commands to build and install binary:

```
go build -o go-inventory
go install
```

<a id="item"><a>
## Item

### Structure

```
type item struct {
  name       string
  brand      string
  sku        string
  category   string
  location   string
}
```
*Stock Keeping Unit (SKU)*

**All *SKU* entries must be unique!**

Below is an example of how item data will be displayed in terminal:

*[AAP12P21] Item: Iphone 12 pro | Brand: Appple | Category: Phone | location: Warehouse*


<a id="usage"></a>
## Usage

 ### Store Item

```
go-inventory store
```

This will prompt users for item data, insert item into database and display inserted item.


### Remove Item

```
go-inventory remove
```

This will prompt users for an item's stock keeping unit (SKU) and remove item if found.

### Update Item

```
go-inventory update
```

This will prompt users for new item data, when entering no new data on prompt the original data stays the same.

### Search Items

```
go-inventory search [optional flag]
```
### Flags
- *name* or *n*
- *brand* or *b*
- *category* or *c*

When no flags are provided the **search** command will prompt user for an item stock keeping unit (SKU).
