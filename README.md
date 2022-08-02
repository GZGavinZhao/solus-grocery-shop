# Solus Grocery Store

Some unofficial 3rd party packages for Solus OS. The idea of setting up such a
repository is inspired by
[prateekmedia's work](https://github.com/prateekmedia/Solus-3rdParty). A big
shout-out to you if you're seeing this :)

This repository is an actual repository (i.e. the packages are built and
uploaded to the cloud where they will get indexed in a cloud function and
uploaded into a storage
bucket), unlike previous attempts that used GitHub as a storage and required
manual indexing. For more info on how this repository is architectured entirely
on the cloud, please see the [Architecture](#architecture) section.

I typically update these packages once a month, and will work on some aesthetic
improvements later.

I just learned how to more formally package using `solbuild`, and will try to do
some more re-builds when I have time.

Notable packages includes:

- cascadia-code
- trezor-bridge
- microsoft-edge
- netease-cloud-music-gtk

## Disclaimer

These packages are **not official**, they are neither supported nor endorsed by
the official Solus devs. Do not ask for help in Solus's (or any other) help
forum, instead create an issue [here](https://gitlab.com/solus-grocery-store/solus-grocery-store/issues).

I do **NOT** have the source code for any of the proprietary software. The
installation files are created by simply decompressing official the .deb
installation file and copying/assigning its contents to the right location(s)
your Solus system (this applies to Microsoft Edge). In fact, in order to install
`microsoft-edge`, you have to convert it yourself. I don't provide the binary
packages.

## Installation Instructions

### General

For users in North America:

```bash
sudo eopkg ar Grocery https://solus-grocery-store-oss.oss-us-east-1.aliyuncs.com/eopkg-index.xml.xz
```

Anywhere else:

```bash
sudo eopkg ar Grocery https://repo-cdn.gzgz.dev/eopkg-index.xml.xz
```

Run `eopkg la Grocery` to see what packages are available.

### Microsoft Edge (stable)

Make sure you properly install `solbuild` first. See [here](https://getsol.us/articles/packaging/building-a-package/en/)
for instructions. You don't need to set up common. There's also no need to build
against unstable.

If you feel like this is too much work, you might be better off using the
[Flatpak](https://discuss.getsol.us/d/6519-microsoft-edge-linux-flatpak)
version. Note that they haven't added stable version yet, so it's only the dev
version.

```bash
git clone https://gitlab.com/solus-grocery-store/solus-grocery-store --depth 1
cd solus-grocery-store/microsoft-edge
sudo solbuild build # This will build/convert the package
sudo eopkg it *.eopkg
```

### Microsoft Edge (dev)

Not really maintained. Use the stable version.

```bash
sudo eopkg bi --ignore-safety https://raw.githubusercontent.com/GZGavinZhao/solus-grocery-store/main/msedge-dev/microsoft-edge-dev/pspec.xml && sudo eopkg it microsoft-edge-dev*.eopkg && sudo rm microsoft-edge-dev*.eopkg
```

**NOTE:** In order to use the latest sign-in and sync support for version
91.0.831.1 and above, after installation you might need to manually type
`edge://flags` in the address bar, then search for and enable the "MSA sign in" experiment.

### 网易云音乐(NetEase Cloud Music)

Not really maintained. Use `netease-cloud-music-gtk` instead.

```
sudo eopkg bi --ignore-safety https://raw.githubusercontent.com/GZGavinZhao/solus-grocery-store/main/netease-cloud-music/ncm/pspec.xml && sudo eopkg it netease-cloud-music-1.2.1-1-1-x86_64.eopkg && sudo rm netease-cloud-music-1.2.1-1-1-x86_64.eopkg
```

**注意事项：** 复制粘贴干就完了。如果启动不了，提交[issue](https://github.com/GZGavinZhao/solus-grocery-store/issues)，同时附上在terminal里面执行`netease-cloud-music`后的所有输出。

**NOTE:** This should work flawlessly, but it's possible that I am missing some
dependencies. If you can't start it, execute `netease-cloud-music` and create
an issue [here](https://github.com/GZGavinZhao/solus-grocery-store/issues) with
the warning/error messages, if any.

### 搜狗拼音(sogoupinyin)

***别用，根本用不了…… 估计依赖有问题
Don't use it, it just doesn't work.***

软件库里已经有Rime输入法，建议用它。
There's already Rime input method in the official repository, which is
recommended over sogoupinyin.

```
sudo eopkg bi --ignore-safety https://raw.githubusercontent.com/GZGavinZhao/solus-grocery-store/main/sogoupinyin/sogoupinyin/pspec.xml && sudo eopkg it sogoupinyin-2.4.0.3469-1-1-x86_64.eopkg && sudo rm sogoupinyin-2.4.0.3469-1-1-x86_64.eopkg
```

**注意事项：** 安装后需要重启。`Fcitx`和`ibus`貌似不能同时使用，在未删掉期间若切换输入法可能会出现系统卡死的现象，唯一的解决办法就是强制重启……所以一定要保存好东西。若无法使用请参考[这篇文章](https://manjaro.org.cn/bbs/topic/manjaro%E4%B8%AD%E6%96%87%E8%BE%93%E5%85%A5%E6%B3%95%EF%BC%88fcitxgooglepinyin%E7%9A%84%E9%85%8D%E7%BD%AE%E9%97%AE%E9%A2%98)

**NOTE:** Reboot after installing. When you use `Fcitx`, it seems that you can't
use any input methods from `ibus` (e.g. the once you set up in system settings).
You can switch to them, but they can only output plain English. If you can't use
it, refer to [this article](https://manjaro.org.cn/bbs/topic/manjaro%E4%B8%AD%E6%96%87%E8%BE%93%E5%85%A5%E6%B3%95%EF%BC%88fcitxgooglepinyin%E7%9A%84%E9%85%8D%E7%BD%AE%E9%97%AE%E9%A2%98).

## Known Issues

### 搜狗拼音(sogoupinyin)

总结：~~能用，但很不稳定~~根本没法用。建议先开`Fcitx`自带的拼音输入法或者`ibus`的`libpinyin`和`Pinyin`以防万一，同时参考安装指南中的文章。

`Fcitx`和`ibus`似乎仍然存在互相干扰的问题……例如，尽管能正常打字/切换中英文，terminal中输入`fcitx`时会显示如下错误：

```
(WARN-7334 dbusstuff.c:248) DBus Service Already Exists
(ERROR-7334 instance.c:443) Exiting.
```

同时，搜狗拼音经常会“诈尸”。有时候它能用，有时候它自己就消失了。还请有能力的大神指点一二！

The interference between `Fcitx` and `ibus` is still a headache for me. It does
work normally, but still gives weird warning messages when running `fcitx` in
the terminal:

```
(WARN-7334 dbusstuff.c:248) DBus Service Already Exists
(ERROR-7334 instance.c:443) Exiting.
```

Please open an issue if you know how might it be resolved. Any help is appreciated!

## Architecture

Alicloud FC+NAS+OSS. WIP.
