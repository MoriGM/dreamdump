Name:           dreamdump
Version:        0.3.0
Release:        %autorelease
Summary:        A command-line fuzzy finder written in Go

License:        MIT
URL:            https://github.com/MoriGM/dreamdump
Source0:        %{expand:%%(pwd)}


%description
Small program to dump dreamcast gdrom discs


%prep
%setup -n dreamdump -c -T
git clone file://%{SOURCEURL0} .


%build
go build -o dreamdump -v .

%install
mkdir -p %{buildroot}/usr/bin
install -Dpm0755 dreamdump %{buildroot}/usr/bin/dreamdump

%files
%license LICENSE
%doc README.md
/usr/bin/dreamdump


%changelog
%autochangelog
