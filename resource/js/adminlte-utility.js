var AdminLteUtility = /** @class */ (function() {
    function AdminLteUtility() {}

    /**
     * ckeditorのエディター生成
     */
    AdminLteUtility.prototype.summernote = function(editorID, options) {
        options = options || {};
        var summernote = $('#' + editorID).summernote(options);
        return summernote;
    };

    /**
     * sort用
     * 親子関係がないソートをする場合はmaxLengthに0を指定する
     */
    AdminLteUtility.prototype.sortEvent = function(submitID, inputName, dataID) {
        inputName = inputName || 'ids';
        dataID = dataID || 'id';

        $submit = $('#' + submitID);

        $submit.on('click', function() {
            var form = $(this).parent('form');
            var elementID = $(this).data('element-id');
            form.children('input[name^="' + inputName + '"]').remove();
            (elementID ? $('#' + elementID + ' [data-' + dataID + ']') : $('[data-' + dataID + ']')).each(function(index, element) {
                var id = $(element).data(dataID);
                form.append($('<input>').val(id).attr('type', 'hidden').attr('name', inputName));
            });
        });
    };

    return AdminLteUtility;
}());
var $Utility = new AdminLteUtility();
window.$Utility = $Utility;