{ pkgs ? import <nixpkgs> { } }:

pkgs.mkShell {
  packages = [
    pkgs.go_1_27
    pkgs.gotools
    pkgs.gopls
  ];

  shellHook = ''
    export GOROOT="${pkgs.go_1_27}/share/go"
    export GOTOOLCHAIN=local

    # Stable path for IDEs (e.g. IntelliJ) to use as the Go SDK location.
    ln -sfn "$GOROOT" .go
  '';
}
