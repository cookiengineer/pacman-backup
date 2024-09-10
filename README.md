
# Pacman Backup

Offline Pacman Cache management tool that allows off-the-grid updates via
sneakernet (usb drives) or mesh Wi-Fi networks.


## Share Updates via USB Drive

In the below example, make sure that `usb-drive` is writable. Replace the path with
the correct one that points to your mounted USB drive.

**Step 1**: On the machine _with_ internet connection, insert and mount the USB drive.

```bash
# Machine with internet connection
pacman-backup download /run/media/$USER/usb-drive;
pacman-backup cleanup /run/media/$USER/usb-drive;
sync;

# Unmount the USB drive and sneak/walk to other machine
```

**Step 2**: On the machine _without_ internet connection, insert and mount the USB drive.

```bash
# Machine without internet connection
pacman-backup upgrade /run/media/$USER/usb-drive;
```

## Share Updates via Wi-Fi Mesh Network

In the below example, the machine _with_ internet connection has the IP `192.168.0.10`.
Replace the IP with the correct one that matches your setup. If in doubt, use `ip` or `ifconfig`.


**Step 1**: On the machine with internet connection, download all updates and serve them as
local pacman archive mirror.

```bash
# Machine with internet connection
sudo pacman-backup download;
pacman-backup serve;
```

**Step 2**: On the machine _without_ direct internet connection, download updates from the
local pacman archive mirror.

```bash
# Machine without internet connection
sudo pacman-backup download http://192.168.0.10:15678/;
sudo pacman-backup upgrade;
```


## Manual Export and Import of Database Files and Package Cache

If you don't trust automated upgrades and want to use `pacman` directly, that's fine. You
can do so by using `export` on the machine with internet connection and `import` on the
machine without internet connection.

```bash
# Machine with internet connection
sudo pacman -Syuw;
pacman-backup export /run/media/$USER/usb-drive;
sync;

# Unmount the USB drive and sneak/walk to other machine
```

```bash
# Machine without internet connection
sudo pacman-backup import /run/media/$USER/usb-drive;
sync;

sudo pacman -Su;
```


# License

GPL3
