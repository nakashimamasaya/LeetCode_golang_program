if [ $# -ne  1 ]; then
  echo "引数の数がおかしいぞ！！！！！"
  exit
fi

mkdir problem/$1
touch problem/$1/$1.go
echo "package _$1" > problem/$1/$1.go
