const canvas = document.getElementById('tutorial');

if (!canvas.getContext) {
    alert('Your browser does not support the canvas element.');
}

const ctx = canvas.getContext('2d');

// Downlods the canvas as an image
function downloadImage() {
    canvas.toBlob(function(blob) {
        const link = document.createElement('a');
        const url = URL.createObjectURL(blob);

        document.body.appendChild(link);
        link.href = url;
        link.download = 'image.png';

        link.click();
        document.body.removeChild(link);
    });
}

document.getElementById('btnDownload').addEventListener('click', downloadImage);
document.getElementById('patternSelector').addEventListener('change', selectPattern)


function selectPattern() {
    const pattern = document.getElementById('patternSelector').value;

    // Clear the canvas
    ctx.clearRect(0, 0, canvas.width, canvas.height);

    switch (pattern) {
        case 'red-blue-square':
            ctx.fillStyle = 'rgb(200, 0, 0)';
            ctx.fillRect(10, 10, 50, 50);
            ctx.fillStyle = 'rgba(0, 0, 200, 0.5)';
            ctx.fillRect(30, 30, 50, 50);
            break;

        case 'black-stoked-square':
            ctx.fillStyle = 'rgb(0, 0, 0)';
            ctx.fillRect(25, 25, 100, 100);
            ctx.clearRect(45, 45, 60, 60);
            ctx.strokeRect(50, 50, 50, 50);
            break;
        case 'checkerboard':
            const squareSize = 10;
            ctx.fillStyle = 'rgb(200, 0, 0)';

            for (let x = 0; x < canvas.width; x += squareSize) {
                for (let y = 0; y < canvas.height; y += squareSize) {
                    if ((x + y) % (squareSize * 2) === 0) {
                        ctx.fillRect(x, y, squareSize, squareSize);
                    }
                }
            }
            break;
        case 'smiley':
            ctx.beginPath();
            ctx.strokeStyle = 'rgb(0, 0, 0)';

            ctx.arc(75, 75, 50, 0, Math.PI * 2, true); // Outer circle
            ctx.moveTo(110, 75);
            ctx.arc(75, 75, 35, 0, Math.PI, false); // Mouth (clockwise)
            ctx.moveTo(65, 65);
            ctx.arc(60, 65, 5, 0, Math.PI * 2, true);
            ctx.moveTo(95, 65);
            ctx.arc(90, 65, 5, 0, Math.PI * 2, true);

            ctx.stroke();
            break;
        case 'triangles':
            ctx.beginPath();
            ctx.fillStyle = 'rgb(0, 0, 0)';
            ctx.strokeStyle = 'rgb(0, 0, 0)';

            ctx.moveTo(25, 25);
            ctx.lineTo(105, 25);
            ctx.lineTo(25, 105);
            ctx.fill();

            ctx.beginPath();
            ctx.moveTo(125, 125);
            ctx.lineTo(125, 45);
            ctx.lineTo(45, 125);
            ctx.closePath();
            ctx.stroke();
    }
}

function init() {
    selectPattern();
}

document.addEventListener("DOMContentLoaded", init)
