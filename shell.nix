with import <nixpkgs> {};

let
  frameworks = darwin.apple_sdk.frameworks;
in mkShell {
  buildInputs = [
    go_1_12
    dep
    autoPatchelfHook
    patchelf
  ];

  shellHook = ''
   GOROOT=${pkgs.go_1_12}/share/go
   export GO111MODULE=on
   export NIX_LDFLAGS="-F${frameworks.WebKit}/Library/Frameworks -framework CoreFoundation $NIX_LDFLAGS";
  '';
}
