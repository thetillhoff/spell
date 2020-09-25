# spell

This repository is about three things:
- A binary called `spell`, which can be used to spell word or sentences in either NATO or german language (see [#usage](#usage)).
- Testing github actions for automated releases
- Getting in touch with installers/packagers like creating a .deb-file and a windows-installer for each release.


## usage
```
$> spell
  usage:
  $> spell [-german|-nato<default>] <what to spell>
$> spell something I cant spell
sierra
oscar
mike
echo
tango
hotel
india
november
golf

india

charlie
alfa
november
tango

sierra
papa
echo
lima
lima
```
Currently, there are two modes:
- `nato`
  is the default but may also be triggered by `spell -nato whatever`. 
- `german`
  can be called likewise with `spell -german wasauchimmer`.
