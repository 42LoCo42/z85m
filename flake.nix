{
  outputs = { flake-utils, nixpkgs, ... }:
    flake-utils.lib.eachDefaultSystem (system:
      let pkgs = import nixpkgs { inherit system; }; in rec {
        packages.default = pkgs.buildGoModule rec {
          pname = "z85m";
          version = "1";
          src = ./.;
          vendorHash = "sha256-X5Zczzl119V23kc1n+qDWsaZIj9UhnqFiZoN6I6s0/Q=";

          CGO_ENABLED = "0";
          ldflags = [ "-s" "-w" ];

          meta.mainProgram = pname;
        };

        devShells.default = pkgs.mkShell {
          inputsFrom = [ packages.default ];
        };
      });
}
