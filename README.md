<div id="top"></div>


[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]



<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/FreezingKas/basic-ransomware-golang">
    <img src="https://i.ibb.co/DfCzBpW/kisspng-malware-computer-icons-computer-virus-virus-5adfcffbea2f32-1822409215246172119592.png" alt="Logo" width="100" height="100">
  </a>

  <h3 align="center">Basic Ransomware GOLANG</h3>

  <p align="center">
    My own ransomware to use my cryptographic knowledges and test GOLANG for the first time
    <br />
    <br />
    <a href="mailto:freezingkas@gmail.com">Contact me</a>
    ·
    <a href="https://github.com/FreezingKas/basic-ransomware-golang/issues">Report Bug</a>
    ·
    <a href="https://github.com/FreezingKas/basic-ransomware-golang/issues">Request Feature</a>
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
        <li><a href="#disclaimer">Disclaimer</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#description">Description</a></li>
        <li><a href="#build">Build</a></li>
      </ul>
    </li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

Recently, I learned cryptography at my university and I wanted to learn GOLANG. I tried to create this ransomware and this is very basic.

Objectives:
* Learn GOLANG syntax.
* Use AES-256 in GOLANG
* Try reversing my own ransomware

<p align="right">(<a href="#top">back to top</a>)</p>

### Disclaimer

Do not use this program for illegal purposes ! I am not responsible for the actions of third parties. Also do not run this program on your machine, use a VM instead.

### Built With

* [Golang](https://go.dev/)
* [AES](https://fr.wikipedia.org/wiki/Advanced_Encryption_Standard/)
* [AES Golang Package](https://pkg.go.dev/crypto/aes)
* [HTTP Golang Package](https://pkg.go.dev/net/http)

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- GETTING STARTED -->
## Getting Started

You have 4 modules : `attacker`, `encrypter`, `decrypter` and `fileutils`. 

`attacker` is the server program to receive the AES key.

`encrypter` is obviously the dangerous program. I repeat, do not launch it on your machine.

`decrypter` decrypt the files with the AES key.

`fileutils` retrieves all the file paths to be encrypted/decrypted. By default it is *C:\Users\\* on Windows and */home/* on Linux. But I redefined the value of **dir** with a testing directory (test_encrypt_dir).

<br>

### Description

* attacker

It is just a simple HTTP server made with the Golang package **net/http**. It receives the AES key in hexadecimal and it stores it in a file named *key.log*.

* encrypter

First, it generates an random 256 bits key and send it to the server. By default it is a POST request made with **net/http** package to localhost:8080 but you can change it to whatever you want in *encrypter.go*. After it retrieves all the content of files in directory defined in *fileutils.go* (you can also change it). Finally it encrypt all the content of files and replace it.

* decrypter

Similarly to encrypter, it retrieves all the files but it decrypt with the key passed in STDIN and replace files contents.

<p align="right">(<a href="#top">back to top</a>)</p>


## Build

First, clone the project :

```
git clone https://github.com/FreezingKas/basic-ransomware-golang.git
```

* Build *attacker.go* :

```sh
cd attacker
go build attacker.go
./attacker
```
It launches the HTTP server on localhost:8080. You can now receive AES keys from *encrypter.go*.

* Build *encrypter.go*
```sh
cd encrypter
go build encrypter.go
./encrypter
```
Obviously, a victim will not build this program itself. After encryption, I have to ask for money but i didn't implement it for the moment.  Don't forget that you can change IP adress.

* Build *decrypter.go*
```sh
cd decrypter
go build decrypter.go
./decrypter
```
The program will ask for the key in command-line and the objectives is to make it simple for a victim to enter the AES key.

<p align="right">(<a href="#top">back to top</a>)</p>

## Roadmap

- [x] Create a server to receive keys.
- [x] Implement function to retrieve files.
- [x] Implement encrypter.
- [x] Implement decrypter.
- [ ] Implement GUI with a threat message and ask for Bitcoin.
- [ ] Implement AES from scratch.
- [ ] Create a database to store keys.

See the [open issues](https://github.com/othneildrew/Best-README-Template/issues) for a full list of proposed features (and known issues).

<p align="right">(<a href="#top">back to top</a>)</p>

## Contributing

Contributions are greatly appreciated, I am still a beginner in Golang. I am not an expert yet. Also I am not a great ransomware creator, so any suggestions are welcomed.

<p align="right">(<a href="#top">back to top</a>)</p>

## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#top">back to top</a>)</p>

## Contact

Your Name - [@Maxence_Jng](https://twitter.com/Maxence_Jng) - freezingkas@gmail.com

Project Link: [https://github.com/FreezingKas/basic-ransomware-golang](https://github.com/FreezingKas/basic-ransomware-golang)

<p align="right">(<a href="#top">back to top</a>)</p>


## Acknowledgments

Use this space to list resources you find helpful and would like to give credit to. I've included a few of my favorites to kick things off!

* [Choose an Open Source License](https://choosealicense.com)
* [Basic README Template](https://github.com/othneildrew/Best-README-Template)
* [Img Shields](https://shields.io)
* [Stackoverflow](https://stackoverflow.com) (as usual)
* [Free PNG](https://www.freepng.fr) (for the logo)

<p align="right">(<a href="#top">back to top</a>)</p>


[contributors-shield]: https://img.shields.io/github/contributors/FreezingKas/basic-ransomware-golang.svg?style=for-the-badge
[contributors-url]: https://github.com/FreezingKas/basic-ransomware-golang/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/FreezingKas/basic-ransomware-golang.svg?style=for-the-badge
[forks-url]: https://github.com/FreezingKas/basic-ransomware-golang/network/members
[stars-shield]: https://img.shields.io/github/stars/FreezingKas/basic-ransomware-golang.svg?style=for-the-badge
[stars-url]: https://github.com/FreezingKas/basic-ransomware-golang/stargazers
[issues-shield]: https://img.shields.io/github/issues/FreezingKas/basic-ransomware-golang.svg?style=for-the-badge
[issues-url]: https://github.com/othneildrew/Best-README-Template/issues
[license-shield]: https://img.shields.io/github/license/FreezingKas/basic-ransomware-golang.svg?style=for-the-badge
[license-url]: https://github.com/FreezingKas/basic-ransomware-golang/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://www.linkedin.com/in/maxence-jung-69501a1a3/
