# Go Inventory

Go Inventory is a tool that allows users to perform simple stock keeping with in the CLI.

## Table of Contents

- [Install](#install)
- [Item](#item)
- [Usage](#usage)

<a id="install"></a>
## Install

To use Go-Inventory a few steps must be taken to set up the command for use.

### Steps
- Create a .env file with your postgres database variables
- install binary file

Create a .env file and add the variables below with your postgres database information:

```
touch .env
```
**Variables**
- *PG_USERNAME*
- *PG_PASSWORD*
- *PG_SERVER*
- *PG_HOST*
- *PG_DATABASE*
- *PG_SSLMODE*

install binary file:

```
go install github.com/LaQuannT/go-inventory/cmd/go-inventory@latest
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
  amount     int
}
```
*Stock Keeping Unit (SKU)*

**All *SKU* entries must be unique!**

Below is an example of how item data will be displayed in terminal:

*[AAP12P21] Name: Iphone 12 pro | Brand: Appple | Category: Phone | location: Warehouse | Amount: 1*


<a id="usage"></a>
## Usage

 ### Store Item

```
go-inventory store
```

This will prompt users for item data, insert item into database and display inserted item.

 ### Add 

```
go-inventory add
```
This will prompt users for item's stock keeping unit and the amount to add to item amount.
negative numbers will subtract from the item's stored amount.

 ### Subtract

```
go-inventory subtract
```
This will prompt users for item's stock keeping unit and the amount to subtract from item amount.

### Delete Item

```
go-inventory delete
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
- *name* or *-n*
- *brand* or *-b*
- *category* or *-c*

When no flags are provided the **search** command will prompt user for an item stock keeping unit (SKU).
