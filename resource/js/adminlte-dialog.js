var DialogUtil = /** @class */ (function() {
    function DialogUtil() {}
    DialogUtil.prototype.can = false;
    DialogUtil.prototype.canSubmit = function() {
        return this.can;
    };
    DialogUtil.prototype.resetSubmit = function() {
        this.can = false;
    };
    DialogUtil.prototype.executeSubmit = function(selector, element) {
        if (selector) {
            $(selector).submit();
        } else {
            this.can = true;
            $(element).click();
        }
    };
    /**
     * data-form=""を設定してFormタグのセレクター指定
     */
    DialogUtil.prototype.delete = function(element) {
        var self = this;

        $('[data-modal="remove"]').remove();

        var selector = $(element).data('form');
        var $btnElement = self.buildButton();
        var $element = self.buildHTML({
            title: 'DELETE',
            message: '削除しますか？',
        });
        $element.find('[data-yes="true"]').on('click', function() {
            self.executeSubmit(selector, element);
        });

        $('body').append($btnElement);
        $('body').append($element);

        $btnElement.trigger('click');
    };
    DialogUtil.prototype.restore = function(element) {
        var self = this;

        $('[data-modal="remove"]').remove();

        var selector = $(element).data('form');
        var $btnElement = self.buildButton();
        var $element = self.buildHTML({
            title: 'RESTORE',
            message: '復元しますか？',
        });
        $element.find('[data-yes="true"]').on('click', function() {
            self.executeSubmit(selector, element);
        });

        $('body').append($btnElement);
        $('body').append($element);

        $btnElement.trigger('click');
    };
    DialogUtil.prototype.update = function(element) {
        var self = this;

        $('[data-modal="remove"]').remove();

        var selector = $(element).data('form');
        var $btnElement = self.buildButton();
        var $element = self.buildHTML({
            title: 'UPDATE',
            message: '更新しますか？',
        });
        $element.find('[data-yes="true"]').on('click', function() {
            self.executeSubmit(selector, element);
        });

        $('body').append($btnElement);
        $('body').append($element);

        $btnElement.trigger('click');
    };
    DialogUtil.prototype.store = function(element) {
        var self = this;

        $('[data-modal="remove"]').remove();

        var selector = $(element).data('form');
        var $btnElement = self.buildButton();
        var $element = self.buildHTML({
            title: 'STORE',
            message: '追加しますか？',
        });
        $element.find('[data-yes="true"]').on('click', function() {
            self.executeSubmit(selector, element);
        });

        $('body').append($btnElement);
        $('body').append($element);

        $btnElement.trigger('click');
    };
    DialogUtil.prototype.message = function(element) {
        var self = this;

        $('[data-modal="remove"]').remove();

        var selector = $(element).data('form');
        var $btnElement = self.buildButton();
        var title = $(element).data('title');
        var message = $(element).data('message');
        var $element = self.buildHTML({
            title: title,
            message: message,
        });
        $element.find('[data-yes="true"]').on('click', function() {
            self.executeSubmit(selector, element);
        });

        $('body').append($btnElement);
        $('body').append($element);

        $btnElement.trigger('click');
    };
    DialogUtil.prototype.messageLink = function(element) {
        var self = this;
        var href = $(element).attr('href');
        $('[data-modal="remove"]').remove();

        var $btnElement = self.buildButton();
        var title = $(element).data('title');
        var message = $(element).data('message');
        var $element = self.buildHTML({
            title: title,
            message: message,
        });
        $element.find('[data-yes="true"]').on('click', function() {
            location.href = href;
        });

        $('body').append($btnElement);
        $('body').append($element);

        $btnElement.trigger('click');
    };
    DialogUtil.prototype.info = function(element) {
        var self = this;

        $('[data-modal="remove"]').remove();

        var title = $(element).data('title');
        var message = $(element).data('message');
        var $element = self.buildHTML({
            title: title,
            message: message,
            no: null,
        });

        $('body').append($element);
    };
    DialogUtil.prototype.instant = function(options) {
        var self = this;

        options = $.extend({
            title: undefined,
            message: undefined,
            yesHandler: function() {
                //
            },
        }, options);

        $('[data-modal="remove"]').remove();

        var $btnElement = self.buildButton();
        var $element = self.buildHTML({
            title: options.title,
            message: options.message,
        });
        $element.find('[data-yes="true"]').on('click', function() {
            options.yesHandler();
        });

        $('body').append($btnElement);
        $('body').append($element);

        $btnElement.trigger('click');
    };
    DialogUtil.prototype.buildButton = function() {
        var html = '<button data-modal="remove" type="button" style="display: none;" data-toggle="modal" data-target="#dialog-modal"></button>';
        return $(html);
    };
    DialogUtil.prototype.buildHTML = function(options) {
        options = $.extend({
            title: 'タイトル',
            message: 'メッセージ',
            yes: 'はい',
            no: 'いいえ',
            class: '',
        }, options);
        var no_html = options.no ? '<button type="button" class="btn btn-secondary btn-outline-light" data-dismiss="modal">' + options.no + '</button>' : '';
        var yes_html = '<button type="button" class="btn btn-primary btn-outline-light" data-dismiss="modal" data-yes="true">' + options.yes + '</button>';
        var html = '<div id="dialog-modal" style="display: none;" class="modal fade" data-modal="remove" style="display: block; padding-right: 17px;" aria-modal="true">' +
            '<div class="modal-dialog">' +
            '<div class="modal-content ' + options.class + '">' +
            '<div class="modal-header">' +
            '<h4 class="modal-title">' + options.title + '</h4>' +
            '<button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">×</span></button>' +
            '</div>' +
            '<div class="modal-body"><p>' + options.message + '</p></div>' +
            '<div class="modal-footer justify-content-between">' +
            no_html +
            yes_html +
            '</div>' +
            '</div>' +
            '</div>' +
            '</div>';
        return $(html);
    };
    return DialogUtil;
}());

var $Dialog = new DialogUtil();
window.$Dialog = $Dialog;
$(function() {
    /**
     * data-class="js-dialog" と data-type="" は必ず設定する
     */
    $('[data-class="js-dialog"]').on('click', function(e) {
        if (!$Dialog.canSubmit()) {
            var type = $(this).data('type');
            switch (type) {
                case 'delete':
                    $Dialog.delete(this)
                    break;
                case 'restore':
                    $Dialog.restore(this)
                    break;
                case 'update':
                    $Dialog.update(this)
                    break;
                case 'store':
                    $Dialog.store(this)
                    break;
                case 'message':
                    $Dialog.message(this)
                    break;
                case 'info':
                    $Dialog.info(this)
                    break;
                default:
            }
            return false;
        }
        $Dialog.resetSubmit();
    });
    /**
     * data-class="js-dialog-link" は必ず設定する
     */
    $('[data-class="js-dialog-link"]').on('click', function(e) {
        var type = $(this).data('type');
        switch (type) {
            case 'message':
            default:
                $Dialog.messageLink(this);
        }
        return false;
    });
});