const dropZone = document.getElementById('dropZone');
const uploadButton = document.getElementById('uploadButton');
const cancelButton = document.getElementById('cancelButton');
const statusMessage = document.getElementById('statusMessage');
const preview = document.getElementById('preview');
let currentFile = null;

function showMessage(message, isError = false) {
    statusMessage.textContent = message;
    statusMessage.className = 'status-message ' + (isError ? 'error' : 'success');
}

function resetUpload() {
    currentFile = null;
    preview.style.display = 'none';
    preview.src = '';
    showMessage('');
    dropZone.querySelector('.upload-text').textContent = 'Перетащите фото сюда';
    dropZone.querySelector('.upload-subtext').style.display = 'block';
}

function handleFile(file) {
    if (!file.type.startsWith('image/')) {
        showMessage('Пожалуйста, выберите изображение', true);
        return;
    }

    currentFile = file;
    const reader = new FileReader();
    reader.onload = (e) => {
        preview.src = e.target.result;
        preview.style.display = 'block';
        dropZone.querySelector('.upload-text').textContent = 'Выбран файл: ' + file.name;
        dropZone.querySelector('.upload-subtext').style.display = 'none';
    };
    reader.readAsDataURL(file);
}

async function uploadFile() {
    if (!currentFile) {
        showMessage('Пожалуйста, выберите файл', true);
        return;
    }

    // Читаем файл как base64
    const reader = new FileReader();
    reader.readAsDataURL(currentFile);
    
    reader.onload = async () => {
        try {
            const response = await fetch('http://localhost:8080/register/image', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    image: reader.result // Отправляем base64 строку
                })
            });

            if (response.ok) {
                showMessage('Файл успешно загружен!');
                setTimeout(resetUpload, 3000);
            } else {
                throw new Error('Ошибка загрузки');
            }
        } catch (error) {
            showMessage('Ошибка при загрузке файла: ' + error.message, true);
        }
    };
}


dropZone.addEventListener('click', () => {
    const input = document.createElement('input');
    input.type = 'file';
    input.accept = 'image/*';
    input.onchange = (e) => {
        const file = e.target.files[0];
        if (file) {
            handleFile(file);
        }
    };
    input.click();
});

['dragenter', 'dragover', 'dragleave', 'drop'].forEach(eventName => {
    dropZone.addEventListener(eventName, preventDefaults, false);
});

function preventDefaults(e) {
    e.preventDefault();
    e.stopPropagation();
}

['dragenter', 'dragover'].forEach(eventName => {
    dropZone.addEventListener(eventName, () => {
        dropZone.style.borderColor = 'rgba(255, 255, 255, 0.8)';
        dropZone.style.background = 'linear-gradient(145deg, rgba(255, 255, 255, 0.1) 0%, rgba(255, 255, 255, 0.15) 100%)';
    });
});

['dragleave', 'drop'].forEach(eventName => {
    dropZone.addEventListener(eventName, () => {
        dropZone.style.borderColor = 'rgba(255, 255, 255, 0.5)';
        dropZone.style.background = 'linear-gradient(145deg, rgba(255, 255, 255, 0.05) 0%, rgba(255, 255, 255, 0.1) 100%)';
    });
});

dropZone.addEventListener('drop', (e) => {
    const file = e.dataTransfer.files[0];
    if (file) {
        handleFile(file);
    }
});

uploadButton.addEventListener('click', uploadFile);
cancelButton.addEventListener('click', resetUpload);