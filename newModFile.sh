#!/bin/bash
cat dict.txt | grep -E '^.{5}$' > 5ltrwd.txt
sed -i -e 's/^/security/' 5ltrwd.txt
perl -p -i -e 'tr/A-Z/a-z/' 5ltrwd.txt