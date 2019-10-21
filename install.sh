echo "Cloning project color-ls"
git clone https://github.com/gabriel-dantas98/klyntar

cd klyntar

echo "Moving binary to /usr/bin/"
sudo mv bin/color-ls /usr/bin/

echo "Nice, you should start using color-ls :D"
color-ls help
