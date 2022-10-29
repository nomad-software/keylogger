# keylogger
**A Linux based keylogger**

---

A simple implementation of a keylogger that works without root access. In fact,
i'm shocked how simple this is to do!

## Usage

```
go run main.go -log /var/log/keylogger.log
```

The specified log will receive all keys pressed while it's executing. There may
be a small time delay when logging keys due to file buffering.

## Dependencies

This will only work on [X11](https://www.x.org) based systems. It will not work
with [Wayland](https://wayland.freedesktop.org).

If your Linux distribution is X11 based then install the following packages to
make sure you have the necessary source files.

```
sudo apt install libx11-dev libxi-dev
```
