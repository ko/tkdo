import Ember from 'ember';

export default Ember.Component.extend({

    actions: {
        createTask: function(task) {
            console.log('kenko:createTask:sendAction()');
            this.sendAction('createTask', task);
            this.set('post', {}); 
        }
    }
});
