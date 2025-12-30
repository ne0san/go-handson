{ pkgs, ... }:

{
  packages = with pkgs; [
    go
    gotools
    gopls
  ];

  languages.go = {
    enable = true;
  };

}
