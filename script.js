//setInterval(colorChange, 50)
/*
var i = 0
var n = 0
var x = 0
///*
                    <span class="logText" style='color: #171738'>F</span>
                    <span class="logText" style='color: #2E1760'>l</span>
                    <span class="logText" style='color: #3423A6'>o</span>
                    <span class="logText" style='color: #7180B9'>w</span>
                    <span class="logText" style='color: #171738'>f</span>
                    <span class="logText" style='color: #2E1760'>l</span>
                    <span class="logText" style='color: #3423A6'>y</span>

function colorChange() {
    var colors = [ /*"#DFF3E4",  "#171738", "#2E1760", "#3423A6", "#7180B9", "#171738", "#2E1760", "#3423A6"];
    if (i > 8) {
        if (x > 8) {
            x = 0;
        } else {
            x++;
        }
        i = x;
    }
    if (n > 8) {
        n = 0;
    }
    $(".logText:eq(" + n + ")").each(function (index, element) {
        console.log("Off " + i + " " + x)
        $(element).css('color', colors[i]);
    });
    i++;
    n++;
    colorChange()
}

//colorChange()

function increase_brightness(hex, percent) {
    // strip the leading # if it's there
    hex = hex.replace(/^\s*#|\s*$/g, '');

    // convert 3 char codes --> 6, e.g. `E0F` --> `EE00FF`
    if (hex.length == 3) {
        hex = hex.replace(/(.)/g, '$1$1');
    }

    var r = parseInt(hex.substr(0, 2), 16),
        g = parseInt(hex.substr(2, 2), 16),
        b = parseInt(hex.substr(4, 2), 16);

    return '#' +
        ((0 | (1 << 8) + r + (256 - r) * percent / 100).toString(16)).substr(1) +
        ((0 | (1 << 8) + g + (256 - g) * percent / 100).toString(16)).substr(1) +
        ((0 | (1 << 8) + b + (256 - b) * percent / 100).toString(16)).substr(1);
}
*/