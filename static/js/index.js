
const editor = new EditorJS({
	holder: 'editorjs',
	logLevel: 'VERBOSE',
	autofocus: true,
	placeholder: 'Welcome to your favorite web editor',
	tools: {
		image: SimpleImage,
	},
	data: {},

});

editor.isReady
	.then(() => {
		console.log("Editorjs is ready for work");

		editor.save().then(savedData => {
			output.innerHTML = JSON.stringify(savedData, null, 4);
		})
	})
	.catch((reason) => console.log(`Editorjs failed to initialize with reason: ${reason}`))

const saveButton = document.getElementById('save-button');
const output = document.getElementById('output');

saveButton.addEventListener('click', () => {
	editor.save().then(savedData => {
		output.innerHTML = JSON.stringify(savedData, null, 4);
	})
});
