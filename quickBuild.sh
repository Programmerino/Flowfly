rm main_temp.js main_temp.js.map  material.js jquery.js output.js output.js.map uglify.js.map
wget -O jquery.js https://code.jquery.com/jquery-3.3.1.slim.min.js
wget -O material.js https://code.getmdl.io/1.3.0/material.min.js
rm main_temp.js
GOOS=linux gopherjs build src/main.go -o main_temp.js
uglifyjs jquery.js main_temp.js material.js --output output.js
#java -jar closure-compiler.jar --compilation_level SIMPLE_OPTIMIZATIONS --js uglify.js --js_output_file output.js --language_in ECMASCRIPT5 --warning_level QUIET --third_party --jscomp_off \* --use_types_for_optimization false
rm main_temp.js jquery.js material.js uglify.js