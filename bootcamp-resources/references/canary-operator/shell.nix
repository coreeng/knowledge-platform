let
  nixpkgs = builtins.fetchTarball {
    name = "nixpkgs-22.11";
    url = "https://github.com/NixOS/nixpkgs/archive/refs/tags/22.11.tar.gz";
    sha256 = "11w3wn2yjhaa5pv20gbfbirvjq6i3m7pqrq2msf0g7cv44vijwgw";
  };

  pkgs = import nixpkgs {
    config = {

    };
  };
  tarballUnstable = builtins.fetchTarball {
    name = "nixpkgs-unstable";
    url = "https://github.com/nixos/nixpkgs/archive/8ad5e8132c5dcf977e308e7bf5517cc6cc0bf7d8.tar.gz";
    sha256 = "17v6wigks04x1d63a2wcd7cc4z9ca6qr0f4xvw1pdw83f8a3c0nj";
  };

pkgsUnstable = import tarballUnstable { };

in pkgs.mkShell {
  buildInputs = with pkgs; [
    gnumake
    docker
    minikube
    pkgsUnstable.go_1_19 # installing go version 1.19.6
    kubectl
  ];
  shellHook = ''
    scripts/install-kubebuilder.sh
  '';
}