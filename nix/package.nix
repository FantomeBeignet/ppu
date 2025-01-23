{
  lib,
  pkgs ? (
    let
      inherit (builtins) fetchTree fromJSON readFile;
      inherit ((fromJSON (readFile ../flake.lock)).nodes) nixpkgs gomod2nix;
    in
      import (fetchTree nixpkgs.locked) {
        overlays = [
          (import "${fetchTree gomod2nix.locked}/overlay.nix")
        ];
      }
  ),
  buildGoApplication ? pkgs.buildGoApplication,
}:
buildGoApplication rec {
  pname = "ppu";
  version = "0.1";
  src = lib.cleanSource ../.;
  modules = ../gomod2nix.toml;
  go = pkgs.go_1_22;
  CGO_ENABLED = 0;
  ldflags = [
    ''-X="git.sr.ht/~fantomebeignet/ppu.version=${version}"''
    "-s"
    "-w"
    ''-extldflags "-static"''
  ];
}
