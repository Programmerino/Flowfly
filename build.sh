rm main_temp.js main_temp.js.map  material.js jquery.js output.js output.js.map uglify.js.map
wget -O jquery.js https://code.jquery.com/jquery-3.3.1.slim.min.js
wget -O material.js https://code.getmdl.io/1.3.0/material.min.js
rm main_temp.js
gopherjs build src/main.go -o main_temp.js
uglifyjs jquery.js main_temp.js material.js -c drop_console,keep_fargs=false,passes=5,toplevel=true,unsafe,unsafe_comps,unsafe_Function,unsafe_math,unsafe_proto,unsafe_regexp,unsafe_undefined --output uglify.js
rm main_temp.js jquery.js material.js
for run in {1..5}
do
java -jar closure-compiler.jar --compilation_level SIMPLE_OPTIMIZATIONS --js uglify.js --js_output_file closure.js --language_in ECMASCRIPT5 --warning_level QUIET --third_party --jscomp_off \* --use_types_for_optimization false
mv closure.js uglify.js
done
mv uglify.js closure.js
rm uglify.js
for run in {1..10}
do
uglifyjs closure.js -c drop_console,keep_fargs=false,passes=1,toplevel=true,unsafe,unsafe_comps,unsafe_Function,unsafe_math,unsafe_proto,unsafe_regexp,unsafe_undefined --output uglify.js
rm closure.js
java -jar closure-compiler.jar --compilation_level SIMPLE_OPTIMIZATIONS --js uglify.js --js_output_file closure.js --language_in ECMASCRIPT5 --warning_level QUIET --third_party --jscomp_off \* --use_types_for_optimization false
rm uglify.js
done
uglifyjs closure.js -c drop_console,keep_fargs=false,passes=1,toplevel=true,unsafe,unsafe_comps,unsafe_Function,unsafe_math,unsafe_proto,unsafe_regexp,unsafe_undefined --output output.js
rm closure.js