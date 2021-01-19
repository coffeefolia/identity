#!/bin/bash
cat dict.txt | grep -E '^.{8}$' > 8ltrwd.txt
sed -i -e 's/^/security/' 8ltrwd.txt
perl -p -i -e 'tr/A-Z/a-z/' 8ltrwd.txt