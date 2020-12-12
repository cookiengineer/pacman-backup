
# Pacman Backup

Backup tool for off-the-grid updates via portable USB sticks or (mesh) LAN networks.


## Share Updates via USB Drive

**Step 1**:

On the machine with internet connection, insert and mount the USB drive.

Use `pacman-backup archive` to copy the package cache to the backup folder.
Use `pacman-backup cleanup` to remain only the latest version of each package.

The `archive` action will also copy the necessary database files for `pacman -Sy`.

In the below example, make sure that `pacman-usbstick` is writable (and replace
accordingly).

```bash
# Machine with internet connection

sudo pacman -Sy;
sudo pacman -Suw;

pacman-backup archive /run/media/$USER/pacman-usbstick;
pacman-backup cleanup /run/media/$USER/pacman-usbstick;
sync;

# Then, unmount the USB drive
```

**Step 2**:

On the machine without internet connection, insert and mount the USB drive.

Copy the necessary database files to pacman's `sync` folder (which is `$DBPath/sync` of `pacman.conf`).

Use `pacman-backup upgrade` to update from the backupfolder.
This will output the pacman command that you should verify manually before executing it.

```bash
# Machine without internet connection

sudo cp /run/media/cookiengineer/pacman-usbstick/sync/*.db /var/lib/pacman/sync/;'
pacman-backup upgrade /run/media/cookiengineer/pacman-usbstick;'
```


## Share Updates via LAN Connection

**Step 1**:

On the machine with internet connection, connect the LAN cable (and internet connection).

Use `pacman-backup serve` to start a pacman server that can be used by other pacman clients.

In the below example, the machine with internet connection has the IP `192.168.0.10`.

```bash
# Machine with internet connection
sudo pacman -Sy;
sudo pacman -Suw;

pacman-backup serve;
```

**Step 2**:

On the machine without internet connection, connect the LAN cable (so that the
server running at `192.168.0.10` is reachable).

Use `pacman-backup download 192.168.0.10` to download the packages to pacman's package cache.

Use `pacman-backup upgrade` to update from the local pacman cache.
This will output the pacman command that you should verify manually before executing it.

```bash
sudo pacman-backup download 192.168.0.10;
pacman-backup upgrade;
```

## Share Updates via LAN Cache Proxy

`pacman-backup` can also emulate a local Cache Proxy for other pacman clients.
If `pacman-backup serve` is running on the machine with internet connection, it
can be used for `pacman` directly.

Note that if the packages don't exist, they will appear in the logs but aren't downloaded
directly; and that partial upgrades are not officially supported by Arch Linux.

**Step 1**:

On the machine with internet connection, connect the LAN cable (and internet connection).

Use `pacman-backup serve` to start a pacman server that can be used by other pacman clients.

In the below example, the machine with internet connection has the IP `192.168.0.10`.

```bash
# Machine with internet connection

sudo pacman -Sy;
sudo pacman -Suw;

pacman-backup serve;
```

**Step 2**:

Modify the `/etc/pacman.d/mirrorlist` to have as a first entry the following line:

```conf
# Pacman Mirrorlist for local server

Server = http://192.168.0.10:15678
```

Use `pacman -Sy` and `pacman -Su` to update from the Cache Proxy.

```bash
sudo pacman -Sy;
sudo pacman -Su; # or use -Suw
```


## Advanced Usage Hints

`pacman-backup` parses the CLI parameters dynamically. If there's a parameter that looks
like a `folder`, the action is executed on the specified folder.

That means it is possible to e.g. run `download`, `serve`, `upgrade`, and others on the
specified folder. The folder parameter is defaulted with `/var/lib/pacman`.

