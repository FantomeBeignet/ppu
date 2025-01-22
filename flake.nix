{
  description = "Nix flake for ppu";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-parts.url = "github:hercules-ci/flake-parts";
    gomod2nix = {
      url = "github:nix-community/gomod2nix";
      inputs.nixpkgs.follows = "nixpkgs";
    };
    treefmt-nix = {
      url = "github:numtide/treefmt-nix";
    };
  };

  outputs = inputs @ {
    flake-parts,
    gomod2nix,
    ...
  }:
    flake-parts.lib.mkFlake {inherit inputs;} {
      imports = [
        inputs.treefmt-nix.flakeModule
      ];
      systems = ["x86_64-linux"];
      perSystem = {
        pkgs,
        config,
        ...
      }: let
        goShell = pkgs.callPackage ./nix/shell.nix {
          inherit (gomod2nix.legacyPackages.${pkgs.system}) mkGoEnv gomod2nix;
        };
      in {
        treefmt.config = {
          projectRootFile = "flake.nix";
          programs = {
            alejandra.enable = true;
            statix.enable = true;
            deadnix.enable = true;
            gofumpt.enable = true;
          };
        };
        packages.default = pkgs.callPackage ./nix/package.nix {
          inherit (gomod2nix.legacyPackages.${pkgs.system}) buildGoApplication;
        };
        devShells.default = pkgs.mkShellNoCC {
          packages = [config.treefmt.build.wrapper];
          inputsFrom = [goShell];
        };
      };
    };
}
