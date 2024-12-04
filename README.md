# pdiddy

`pdiddy` is a command-line tool that allows you to manage a virtual party while tracking the use of babyoil and system users (guests). It keeps track of which users have been babyoiled, how much babyoil is left in the house, and lets you add more babyoil as needed.

![P -Diddy-1-2448788980](https://github.com/user-attachments/assets/5f917a04-dbab-4995-98f7-e13720a24f5b)



### Features

- **Guestlist**: View all system users on the machine
- **Party**: See currently logged-in users.
- **Babyoil Users**: Track and apply babyoil to individual users.
- **Reset Babyoil Usage**: Reset babyoil usage for a specific user.
- **Check Babyoil**: Display how much babyoil remains in the house.
- **Buy Babyoil**: Add more babyoil to the house.

### Installation

1. Clone this repository or download the source code.
2. Build the programm using Go:
```bash
go build -o pdiddy
```

3. Place the `pdiddy` binary in a directory in your `PATH` for easier use, such as `usr/local/bin`.


## Usage

Run `pdiddy <command>` with one of the following commands:

1. `guestlist`
List all system users on the machine.

```bash
pdiddy guestlist
```

2. `party`
List all currently logged-in users (party attendees).
```bash
pdiddy party
```

3. `babyoil <user>`
Apply babyoil to a user. This reduces the total babyoil in the house by 1 and keeps track of how many times the user has been babyoiled.
```bash
pdiddy babyoil <user>
```

4. `touch <user`
Reset the babyoil usage for a specific user.
```bash
pdiddy touch <user>
```

5. `babyoilhouse`
Display the remaining amount of babyoil in the house
```bash
pdiddy babyoilhouse
```
6. `buy <amount>`
Buy more babyoil. Specify the amount to add to the house
```bash
pdiddy buy <amount>
```

7. `help`
Show the help menu with availble commands and their descriptions
```bash
pdiddy help
```
