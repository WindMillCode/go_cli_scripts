
# FileName: \.vscode\settings.json 

{}

# FileName: \.vscode\tasks.json 


# FileName: \content.js 


# FileName: \index.html 

<html>
<head>
  <title>Modify ChatGPT Prompt</title>
  <style>
    input,button{
      background-color:lightcoral;

    }
    button{
      border-radius: calc(50/16 * 1em);
    }
    input::placeholder{
      color: black;

    }
    body{
      background-color: #171a1b;
      color:white;

    }
    .updateGroupName{
      display: flex;
      justify-content: space-between;
    }
    .buttons0{
      margin: calc(10/16 * 1em) 0 0 0;
      display: flex;
      justify-content: space-between;
    }

    .buttons1{
      height: calc(50/16 * 1em);
      display: flex;
      flex-direction: column;
      justify-content: space-between;
    }

  </style>
  <script src="popup.js"></script>
</head>
<body>
  <h1>Modify Prompt</h1>

  <label for="groupSelect">Select Group:</label>
  <select id="groupSelect"></select><br>
  <div class="updateGroupName">
    <input type="text" id="newGroupName" placeholder="New Group Name">
    <button id="updateGroupName">Update Group Name</button><br>
  </div>
  <label for="prefix">Prefix:</label><br>
  <textarea id="prefix" rows="5" cols="50"></textarea><br>
  <label for="suffix">Suffix:</label><br>
  <textarea id="suffix" rows="5" cols="50"></textarea><br>

  <label for="prefixTimes">Prefix Times:</label><br>
  <input type="number" id="prefixTimes" value="1"><br>
  <label for="suffixTimes">Suffix Times:</label><br>
  <input type="number" id="suffixTimes" value="1"><br>

  <div class="buttons0">
    <div class="buttons1">
      <button id="addGroup">Add Group</button>
      <button id="deleteGroup">Delete Group</button>
    </div>
    <div class="buttons1">
      <button id="importGroups">Import Groups</button>
      <button id="exportGroups">Export Groups</button>

    </div>
  </div>
  <div class="buttons0">
    <div class="buttons1">
      <button id="saveGroup">Save Current Group</button>
    </div>
    <div class="buttons1">
      <button id="modify">Modify</button>
      <button id="undo">Undo</button>
    </div>
  </div>




</body>
</html>

# FileName: \manifest.json 

{
  "manifest_version": 3,
  "name": "ChatGPT Prompt Modifier",
  "version": "1.0",
  "permissions": ["activeTab", "storage", "scripting"],
  "action": {
    "default_popup": "index.html",
    "default_icon": "icon.png"
  },
  "content_security_policy": {
    "extension_pages": "script-src 'self'; object-src 'self'; img-src 'self' blob: data:"
  },
  "content_scripts": [
    {
      "matches": ["<all_urls>"],
      "js": ["content.js"]
    }
  ]
}

# FileName: \popup.js 

function debounce(func, wait) {
  let timeout;
  return function (...args) {
    clearTimeout(timeout);
    timeout = setTimeout(() => func.apply(this, args), wait);
  };
}

async function setStorageItem(key, value) {
  return new Promise((resolve, reject) => {
    chrome.storage.sync.set({ [key]: value }, function() {
      if (chrome.runtime.lastError) {
        reject(chrome.runtime.lastError);
      } else {
        resolve();
      }
    });
  });
}

async function getStorageItem(key) {
  return new Promise((resolve, reject) => {
    chrome.storage.sync.get([key], function(result) {
      if (chrome.runtime.lastError) {
        reject(chrome.runtime.lastError);
      } else {
        resolve(result[key] !== undefined ? result[key] : null);
      }
    });
  });
}


document.addEventListener('DOMContentLoaded',async () => {

  chrome.tabs.query({ active: true, currentWindow: true }, (tabs) => {
    chrome.scripting.executeScript({
      target: { tabId: tabs[0].id },
      func: (prefix, suffix) => {
        const inputField = document.querySelector('textarea'); // Adjust selector based on ChatGPT's input field

      },
      args: [prefix.value, suffix.value]
    });
  });



  const prefix = document.getElementById('prefix');
  const suffix = document.getElementById('suffix');
  const prefixTimes =document.getElementById('prefixTimes')
  const suffixTimes =document.getElementById('suffixTimes')
  const groupSelect = document.getElementById('groupSelect');



  const groups = JSON.parse(await getStorageItem ('groups')) || {};
  populateSelect(groupSelect, groups);

  // Load the selected group
  if (groupSelect.value) {
    const selectedGroup = groups[groupSelect.value];
    if (selectedGroup) {
      prefix.value = selectedGroup.prefix;
      suffix.value = selectedGroup.suffix;
    }
  }


  function populateSelect(selectElement, groups) {
    selectElement.innerHTML = '';
    for (const key in groups) {
      const option = document.createElement('option');
      option.value = key;
      option.text = key;
      selectElement.add(option);
    }
  }

  const saveGroup = debounce(async() => {
    const groupName = groupSelect.value || 'default';
    const groups = JSON.parse(await getStorageItem ('groups')) || {};
    groups[groupName] = {
      prefix: prefix.value,
      suffix: suffix.value,
      prefixTimes: parseInt(prefixTimes.value, 10), // Add this line
      suffixTimes: parseInt(suffixTimes.value, 10)  // Add this line
    };
    await setStorageItem('groups', JSON.stringify(groups));
    await setStorageItem('lastSelectedGroup', groupName);
  }, 300);



  prefix.addEventListener("input", saveGroup);
  suffix.addEventListener("input", saveGroup);
  prefixTimes.addEventListener('input', saveGroup);
  suffixTimes.addEventListener('input', saveGroup);

  document.getElementById('saveGroup').addEventListener('click',async () => {
    const groupName = groupSelect.value || 'default';
    const storedGroups = JSON.parse(await getStorageItem ('groups')) || {};
    storedGroups[groupName] = { prefix: prefix.value, suffix: suffix.value };
    await setStorageItem('groups', JSON.stringify(storedGroups));
  });

  document.getElementById('addGroup').addEventListener('click', async() => {
    const groupName = prompt('Enter a name for this group:');
    if (groupName) {
      const storedGroups = JSON.parse(await getStorageItem ('groups')) || {};
      storedGroups[groupName] = { prefix: prefix.value, suffix: suffix.value };
      await setStorageItem('groups', JSON.stringify(storedGroups));
      populateSelect(groupSelect, storedGroups);
      groupSelect.value = groupName;
    }
  });

  document.getElementById('deleteGroup').addEventListener('click', async() => {
    const groupName = groupSelect.value;
    if (groupName && confirm('Are you sure you want to delete this group?')) {
      const storedGroups = JSON.parse(await getStorageItem ('groups')) || {};
      delete storedGroups[groupName];
      await setStorageItem('groups', JSON.stringify(storedGroups));
      populateSelect(groupSelect, storedGroups);
      if (groupSelect.value) {
        const selectedGroup = storedGroups[groupSelect.value];
        prefix.value = selectedGroup ? selectedGroup.prefix : '';
        suffix.value = selectedGroup ? selectedGroup.suffix : '';
      } else {
        prefix.value = '';
        suffix.value = '';
      }
    }
  });

  groupSelect.addEventListener('change', async () => {
    await setStorageItem('lastSelectedGroup', groupSelect.value);
    onGroupChange(groupSelect, prefix, suffix)();
  });


  document.getElementById('modify').addEventListener('click',async () => {
    const prefixValue = prefix.value;
    const suffixValue = suffix.value;
    const prefixTimes = parseInt(document.getElementById('prefixTimes').value, 10);
    const suffixTimes = parseInt(document.getElementById('suffixTimes').value, 10);

    const groupName = groupSelect.value || 'default';
    const storedGroups = JSON.parse(await getStorageItem ('groups')) || {};
    storedGroups[groupName] = { prefix: prefixValue, suffix: suffixValue };
    await setStorageItem('groups', JSON.stringify(storedGroups));

    chrome.tabs.query({ active: true, currentWindow: true }, (tabs) => {
      chrome.scripting.executeScript({
        target: { tabId: tabs[0].id },
        func: (prefix, suffix, prefixTimes, suffixTimes) => {
          const inputField = document.querySelector('#prompt-textarea');
          if (inputField) {
            let modifiedPrompt = inputField.value.trim();
            for (let i = 0; i < prefixTimes; i++) {
              modifiedPrompt = `${prefix} ${modifiedPrompt}`;
            }
            for (let i = 0; i < suffixTimes; i++) {
              modifiedPrompt = `${modifiedPrompt} ${suffix}`;
            }
            inputField.value = modifiedPrompt;
            inputField.dispatchEvent(new Event('input', { bubbles: true }));
            inputField.dispatchEvent(new Event('change', { bubbles: true }));
          }
        },
        args: [prefixValue, suffixValue, prefixTimes, suffixTimes]
      });
    });
  });

  document.getElementById('importGroups').addEventListener('click',async () => {
    const fileInput = document.createElement('input');
    fileInput.type = 'file';
    fileInput.accept = 'application/json';
    fileInput.onchange = (event) => {
      const file = event.target.files[0];
      const reader = new FileReader();
      reader.onload = async (e) => {
        const groups = JSON.parse(e.target.result);
        await setStorageItem('groups', JSON.stringify(groups));
        populateSelect(groupSelect, groups);
        onGroupChange(groupSelect, prefix, suffix)()
      };
      reader.readAsText(file);
    };
    fileInput.click();
  });

  document.getElementById('exportGroups').addEventListener('click',async () => {
    const groups = JSON.parse(await getStorageItem ('groups')) || {};
    const blob = new Blob([JSON.stringify(groups, null, 2)], { type: 'application/json' });
    const url = URL.createObjectURL(blob);
    const a = document.createElement('a');
    a.href = url;
    a.download = 'groups.json';
    a.click();
    URL.revokeObjectURL(url);
  });

  let originalPrompt = '';
  chrome.tabs.query({ active: true, currentWindow: true }, (tabs) => {
    chrome.scripting.executeScript({
      target: { tabId: tabs[0].id },
      func: () => document.querySelector('#prompt-textarea')?.value?.trim?.() ?? "",
    }, (result) => {
      originalPrompt = result[0].result || '';
    });
  });

  document.getElementById('undo').addEventListener('click', () => {
    chrome.tabs.query({ active: true, currentWindow: true }, (tabs) => {
      chrome.scripting.executeScript({
        target: { tabId: tabs[0].id },
        func: (originalPrompt) => {
          const inputField = document.querySelector('#prompt-textarea');
          if (inputField) {
            inputField.value = originalPrompt;
            inputField.dispatchEvent(new Event('input', { bubbles: true }));
            inputField.dispatchEvent(new Event('change', { bubbles: true }));
          }
        },
        args: [originalPrompt]

      });
    });
  });



  document.getElementById('updateGroupName').addEventListener('click', async () => {
    const oldGroupName = groupSelect.value;
    const newGroupName = document.getElementById('newGroupName').value.trim();
    if (oldGroupName && newGroupName) {
      const storedGroups = JSON.parse(await getStorageItem ('groups')) || {};
      if (storedGroups[oldGroupName]) {
        storedGroups[newGroupName] = storedGroups[oldGroupName];
        delete storedGroups[oldGroupName];
        await setStorageItem('groups', JSON.stringify(storedGroups));
        populateSelect(groupSelect, storedGroups);
        groupSelect.value = newGroupName;
        document.getElementById('newGroupName').value = '';
      }
    }
  });



  const lastSelectedGroup = await getStorageItem ('lastSelectedGroup');
  if (lastSelectedGroup && groups[lastSelectedGroup]) {
    const selectedGroup = groups[lastSelectedGroup];
    groupSelect.value = lastSelectedGroup;
    prefix.value = selectedGroup.prefix;
    suffix.value = selectedGroup.suffix;
    prefixTimes.value = selectedGroup.prefixTimes || 1;
    suffixTimes.value = selectedGroup.suffixTimes || 1;
  }

});

function onGroupChange(groupSelect, prefix, suffix) {
  return async () => {
    const groups = JSON.parse(await getStorageItem ('groups')) || {};
    const selectedGroup = groups[groupSelect.value];
    if (selectedGroup) {
      prefix.value = selectedGroup.prefix;
      suffix.value = selectedGroup.suffix;
      document.getElementById('prefixTimes').value = selectedGroup.prefixTimes || 1;
      document.getElementById('suffixTimes').value = selectedGroup.suffixTimes || 1;
    }
  };
}

function modifyChatGPTPrompt(prefix, suffix) {
  // Find the ChatGPT input field
  const inputField = document.querySelector('textarea'); // Adjust selector based on ChatGPT's input field
  if (inputField) {
    // Prepend and append prefix and suffix to the current input field value
    inputField.value = `${prefix} ${inputField.value.trim()} ${suffix}`;
  }
}
