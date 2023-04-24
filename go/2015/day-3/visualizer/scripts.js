const canvas = document.getElementById('visualizer');
const ctx = canvas.getContext('2d');
const scale = 30;

function init() {
    listFiles();

}

const patternSelector = document.getElementById('patternSelector');

document.addEventListener('DOMContentLoaded', init);
patternSelector.addEventListener('change', loadData);

/**
 * List all files in the data directory
 */
async function listFiles() {
    let files = [];

    const data = await fetch('http://localhost:3000/data')
        .then(async response => {
            const data = await response.json();

            // Clear the select element
            patternSelector.innerHTML = '`<option value="">Select a file</option>`;';

            // file: "{name: 'sample.txt'}"
            data.files.forEach(file => {
                patternSelector.innerHTML += `<option value="${file.name}">${file.name}</option>`;
            })
        })
        .catch(error => { console.log(error) })
}

async function loadData() {
    const fileName = patternSelector.value;

    if (fileName === '')
        return;

    const data = await fetch(`http://localhost:3000/data/${fileName}`)
        .then(async response => {
            const data = await response.json();
            const size = resizeCanvas(data.steps);

            createGrid(size);
            plotPath(data.steps, size)
        })
        .catch(error => { console.log(error) })
}

// Get the size of the canvas, based on the coodinates
function resizeCanvas(steps) {
    // Syntax:
    // "steps": [{"coords":{"x": 0, "y": 0}, "entity": "R", "direction": "^"}]
    let size = {
        minX: 0,
        maxX: 0,
        minY: 0,
        maxY: 0,
        height: 0,
        width: 0,
    }

    steps.forEach(step => {
        if (step.coords.x < size.minX) {
            size.minX = step.coords.x;
        }

        if (step.coords.x > size.maxX) {
            size.maxX = step.coords.x;
        }

        if (step.coords.y < size.minY) {
            size.minY = step.coords.y;
        }

        if (step.coords.y > size.maxY) {
            size.maxY = step.coords.y;
        }
    });

    // Resize the canvas
    let height = Math.abs(size.minY) + size.maxY;
    let width = Math.abs(size.minX) + size.maxX;

    size.height = height;
    size.width = width;

    canvas.width = (width * scale) + scale;
    canvas.height = (height * scale) + scale;

    return size;
}

function createGrid(size) {
    // Clear the canvas
    ctx.fillStyle = '#FFF';
    ctx.clearRect(0, 0, canvas.width, canvas.height);
    
    for (let x = scale; x <= canvas.width - scale; x += scale) {
        for (let y = scale; y <= canvas.height - scale; y += scale) {
            // Create the grid with circles
            ctx.fillStyle = '#0087D7';
            ctx.beginPath();
            ctx.arc(x, y, 2, 0, 2 * Math.PI);
            ctx.fill();
            ctx.closePath();
        }
    }
}

// Plot the path of the robot (R - orange) and santa (S - red)
function plotPath(steps, size) {
    const entitySize = 5;

    const robotColor = 'rgba(255, 165, 0, 0.5)';
    const santaColor = 'rgba(255, 0, 123, 0.5)'
    const axisColor = 'rgba(0, 0, 0, 0.5)';

    let robotPos = {
        x: 0,
        y: 0,
        previousX: 0,
        previousY: 0
    }

    let santaPos = {
        x: 0,
        y: 0,
        previousX: 0,
        previousY: 0
    }

    // Plot the origin
    const centerX = (Math.abs(size.minX) * scale) + scale;
    const centerY = (Math.abs(size.minY) * scale) + scale;

    ctx.fillStyle = axisColor;
    ctx.beginPath();
    ctx.lineWidth = 2;
    ctx.arc(centerX, centerY, entitySize, 0, 2 * Math.PI);
    ctx.fill();
    ctx.closePath();

    // Plot the axis

    // y axis
    ctx.strokeStyle = axisColor;
    ctx.beginPath();
    ctx.moveTo(centerX, 0);
    ctx.lineTo(centerX, canvas.height);
    ctx.stroke();
    ctx.closePath();

    // x axis
    ctx.beginPath();
    ctx.moveTo(0, centerY);
    ctx.lineTo(canvas.width, centerY);
    ctx.stroke();

    // Plot the robot's path

    let angle = 0;
    let entity = null;

    for (let i = 0; i < steps.length; i++) {
        x = (steps[i].coords.x + Math.abs(size.minX)) * scale + scale;
        y = (-steps[i].coords.y + Math.abs(size.minY)) * scale + scale;

        ctx.strokeStyle = steps[i].entity == 'R' ? robotColor : santaColor;
        ctx.beginPath();
        ctx.lineWidth = 2;
        ctx.arc(x, y, entitySize, 0, 2 * Math.PI);
        ctx.stroke();
        ctx.closePath();
    }
    
    // for (let i = 0; i < steps.length; i++) {
    //     if (steps[i].entity == "R") continue;
    //     console.log(steps[i]);
    //     ctx.strokeStyle = steps[i].entity == "R" ? robotColor : santaColor;

    //     entity = steps[i].entity == "R" ? robotPos : santaPos;

    //     ctx.beginPath();
    //     ctx.lineWidth = 2;

    //     entity.previousX = entity.x;
    //     entity.previousY = entity.y;

    //     entity.x = (steps[i].coords.x + Math.abs(size.minX)) * scale + scale;
    //     entity.y = (steps[i].coords.y + Math.abs(size.minY)) * scale + scale;

    //     // ctx.moveTo(entity.previousX, entity.previousY);
    //     // ctx.lineTo(entity.x, entity.y);
    //     // ctx.stroke();
    //     // ctx.closePath();

    //     if (steps[i].entity == "R") {
    //         robotPos = entity;
    //     } else {
    //         santaPos = entity;
    //     }
    // }
}
