if [ $# -ne  3 ]; then
  echo "引数の数がおかしいぞ！！！！！"
  exit
fi

mkdir problem/$1
touch problem/$1/$1.go
echo "package _$1\n// $2\n// $3" > problem/$1/$1.go

touch problem/$1/$1_test.go
echo "package _$1" > problem/$1/$1_test.go
