import Ember from 'ember';

export default Ember.Component.extend({
    actions: {

        createTask: function(task) {
            this.sendAction('createTask', task);
            this.set('task', {});
        },

        completeTask: function(task) {
            this.sendAction('completeTask', task);
            this.set('task', {});
        }
    }
});
