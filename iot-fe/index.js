import makePutRequest from "./request.js";

var column1 = document.querySelector(".column-1");
var column2 = document.querySelector(".column-2");
var Pvalue = document.querySelector(".power-value");
var Uvalue = document.querySelector(".u-value");
var Ivalue = document.querySelector(".i-value");
var column1Active = document.getElementById("c1");
var column2Active = document.getElementById("c2");
var Numbervalue = document.querySelector(".number-value");
var SwitchBtn = document.querySelector(".switch-btn");
var activateArr = [];
var deactivateArr = [];
var indexNow;
const url = "http://127.0.0.1:8080/";
fetch(url)
  .then((res) => res.json())
  .then((data) => {
    App(data.objects);
  });

function FilterData(data) {
  data.forEach((object) => {
    if (object.Status == true) {
      activateArr.push(object);
    } else {
      deactivateArr.push(object);
    }
  });
  RenderColumn();
}

function RenderColumn() {
  var renderColumn1 = activateArr.map((object) => {
    return `<p class = "activate">Node ${object.Node_id}</p>`;
  });
  column1.innerHTML = renderColumn1.join("");

  var renderColumn2 = deactivateArr.map((object) => {
    return `<p class = "deactivate">Node ${object.Node_id}</p>`;
  });
  column2.innerHTML = renderColumn2.join("");
}

function HandleBtn() {
  SwitchBtn.onclick = () => {
    if (indexNow == null) {
      SwitchBtn.checked = true;
      alert("Please choose 1 node");
    } else {
      if (activateArr.length >= deactivateArr.length) {
        if (activateArr[indexNow].Status && !SwitchBtn.checked) {
          activateArr[indexNow].Status = false;
          var newData = { status: false };
          makePatchRequest(`${url + activateArr[indexNow].Node_id}`, newData);
          deactivateArr.push(activateArr[indexNow]);
          activateArr.splice(indexNow, 1);
        } else {
          deactivateArr[indexNow].Status = true;
          var newData = { status: true };
          makePatchRequest(`${url + activateArr[indexNow].Node_id}`, newData);
          activateArr.push(deactivateArr[indexNow]);
          deactivateArr.splice(indexNow, 1);
        }
      } else {
        if (!deactivateArr[indexNow].Status && SwitchBtn.checked) {
          deactivateArr[indexNow].Status = true;
          var newData = { status: true };
          makePatchRequest(`${url + activateArr[indexNow].Node_id}`, newData);
          activateArr.push(deactivateArr[indexNow]);
          deactivateArr.splice(indexNow, 1);
        } else {
          activateArr[indexNow].Status = false;
          var newData = { status: false };
          makePatchRequest(`${url + activateArr[indexNow].Node_id}`, newData);
          deactivateArr.push(activateArr[indexNow]);
          activateArr.splice(indexNow, 1);
        }
      }

      indexNow = null;
      RenderColumn();
      DisplayInformation();
    }
  };
}

function DisplayInformation() {
  var ActivateNodes = document.querySelectorAll(".activate");
  ActivateNodes.forEach((node, index) => {
    node.onclick = function () {
      Pvalue.innerText = `${activateArr[index].Power}W`;
      Uvalue.innerText = `${activateArr[index].Voltage}V`;
      Ivalue.innerText = `${activateArr[index].Current}A`;
      Numbervalue.innerText = `${activateArr[index].Energy}`;
      SwitchBtn.checked = true;
      indexNow = index;
    };
  });

  var DeactivateNodes = document.querySelectorAll(".deactivate");
  DeactivateNodes.forEach((node, index) => {
    node.onclick = function () {
      Pvalue.innerText = `${deactivateArr[index].Power}W`;
      Uvalue.innerText = `${deactivateArr[index].Voltage}V`;
      Ivalue.innerText = `${deactivateArr[index].Current}A`;
      Numbervalue.innerText = `${deactivateArr[index].Energy}`;
      SwitchBtn.checked = false;
      indexNow = index;
    };
  });
}

function App(data) {
  FilterData(data);
  DisplayInformation();
  HandleBtn();
}
