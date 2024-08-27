{
  description = "Useful flakes for golang and Kubernetes projects";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = inputs @ { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      with nixpkgs.legacyPackages.${system}; rec {
        packages = rec {
          golangci-lint = pkgs.golangci-lint.override { buildGoModule = buildGo121Module; };

          yamllint-checkstyle = buildGo121Module {
            pname = "yamllint-checkstyle";
            name = "yamllint-checkstyle";
            src = fetchFromGitHub {
              owner = "thomaspoignant";
              repo = "yamllint-checkstyle";
              rev = "v1.0.2";
              sha256 = "jdgzR+q7IiEpZid0/L6rtkKD8d6DvN48rfJZ+EN+xB0=";
            };
            vendorHash = "sha256-LHRd8Q/v3ceFOqULsTtphfd4xBsz3XBG4Rkmn3Ty6CE=";
          };

          keploy = buildGo121Module {
            name = "keploy-server";
            src = fetchFromGitHub {
              owner = "keploy";
              repo = "keploy";
              rev = "v0.9.1";
              hash = "sha256-rIXydYEaVEhhp8J0CbFpdjOCe3O3Yo+2BSUObHUKOB8=";
            };
            doCheck = false;
            subPackages = [ "cmd/server" ];
            vendorHash = "sha256-daLoymQ2Azw5zUHQ9hR9noEQoJJXl/XPOzlOte7RAFE=";
            GOFLAGS = "-mod=readonly";
            patches = [
              # Without this, we get error messages like:
              # vendor/golang.org/x/sys/unix/mremap.go:41:10: unsafe.Slice requires go1.17 or later (-lang was set to go1.16; check go.mod)
              # The patch was generated by changing "go 1.16" to "go 1.21" and executing `go mod tidy`.
              ./keploy/fix-go-version-error.patch
            ];
            buildPhase = ''
                go build -o keploy-server ./cmd/server
            '';
            installPhase = ''
                mkdir -p $out/bin
                cp keploy-server $out/bin/keploy-server
            '';
          };
        };

        formatter = alejandra;
      }
    );
}
