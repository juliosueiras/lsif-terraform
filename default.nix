with import <nixpkgs> {};

buildGoModule rec {
  name = "lsif-terraform";
  version = "0.0.1";
  src = ./.;

  modSha256 = "1mb3169vdlv4h10k15pg88s48s2b6y7v5frk9j9ahg52grygcqb2"; 
  
  goPackagePath = "github.com/juliosueiras/lsif-terraform";
  subPackages = [ "." ];

  meta = with stdenv.lib; {
    description = "Language Server Index Format for Terraform";
    homepage = https://github.com/juliosueiras/terraform-lsp;
    license = licenses.mit;
    maintainers = with maintainers; [ juliosueiras ];
    platforms = platforms.all;
  };
}
