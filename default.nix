with import <nixpkgs> {};

buildGoModule rec {
  name = "lsif-terraform";
  version = "0.0.1";
  src = ./.;

  modSha256 = "1v6ibippl2f6cw45l6dlsid7k8skxn4kcyv25qdj848i83546638"; 
  
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
